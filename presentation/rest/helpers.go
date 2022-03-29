package rest

import (
	"crypto/sha256"
	"crypto/subtle"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"regexp"
	"strings"

	"github.com/VicOsewe/Order-service/domain"
	"github.com/ttacon/libphonenumber"
)

func ValidateCustomerInfo(customer domain.Customer) error {
	if customer.FirstName == "" || customer.LastName == "" || customer.PhoneNumber == "" || customer.Password == "" {
		return fmt.Errorf("invalid  request data, ensure firstname, lastname, phone_number, password is provided")
	}
	err := ValidateEmail(customer.Email)
	if err != nil {
		return err
	}
	err = ValidatePhoneNumber(customer.PhoneNumber)
	if err != nil {
		return err
	}
	return nil
}

func ValidateEmail(email string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return fmt.Errorf("invalid email address")
	}
	return nil
}

func ValidatePhoneNumber(phoneNumber string) error {
	if len(phoneNumber) < 10 {
		return fmt.Errorf("invalid phone number")
	}
	re := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	if !re.MatchString(phoneNumber) {
		return fmt.Errorf("invalid phone number")
	}

	return nil
}

func NormalizePhoneNumber(phoneNumber string) (string, error) {
	defaultRegion := "KE"
	num, err := libphonenumber.Parse(phoneNumber, defaultRegion)
	if err != nil {
		return "", err
	}
	formatted := libphonenumber.Format(num, libphonenumber.INTERNATIONAL)
	cleaned := strings.ReplaceAll(formatted, " ", "")
	cleaned = strings.ReplaceAll(cleaned, "-", "")
	return cleaned, nil
}

func UnmarshalJSONToStruct(w http.ResponseWriter, r *http.Request, targetStruct interface{}) error {
	err := json.NewDecoder(r.Body).Decode(targetStruct)
	if err != nil {
		return err
	}
	return nil
}

func HandlerResponse(w http.ResponseWriter, code int, payload interface{}) {
	marshalledPayload, err := json.Marshal(payload)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	_, err = w.Write(marshalledPayload)
	if err != nil {
		log.Printf(
			"unable to write payload `%s` to the http.ResponseWriter: %s",
			string(marshalledPayload),
			err,
		)
	}
}

func (h *HandlersImplementation) BasicAuth() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				username, password, ok := r.BasicAuth()
				if ok {
					usernameHash := sha256.Sum256([]byte(username))
					passwordHash := sha256.Sum256([]byte(password))
					expectedUsernameHash := sha256.Sum256([]byte(h.auth.username))
					expectedPasswordHash := sha256.Sum256([]byte(h.auth.password))

					usernameMatch := (subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1)
					passwordMatch := (subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1)

					if usernameMatch && passwordMatch {
						next.ServeHTTP(w, r)
						return
					}
				}

				w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
			},
		)
	}
}
