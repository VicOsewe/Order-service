package ait

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/VicOsewe/Order-service/application"
	"github.com/VicOsewe/Order-service/domain/dto"
)

// ServiceAITImpl represents AIT usecases
type ServiceAITImpl struct {
	client http.Client
}

func New(client http.Client) ServiceAITImpl {
	return ServiceAITImpl{client: client}
}

// NewAITService initializes a new instance of the service
func NewAITService() *ServiceAITImpl {
	return &ServiceAITImpl{
		client: http.Client{},
	}
}

func (s *ServiceAITImpl) SendSMS(message, phoneNumber string) error {
	aitUrl := application.GetEnv("AIT_URL")
	apiKey := application.GetEnv("AIT_API_KEY")
	username := application.GetEnv("AIT_USERNAME")
	values := url.Values{}

	values.Set("username", username)
	values.Set("message", message)
	values.Set("to", phoneNumber)
	response, err := s.NewRequest(aitUrl, apiKey, http.MethodPost, values)
	if err != nil {
		return err
	}

	aitResponse := dto.SMSResponse{}

	defer response.Body.Close()
	d, _ := ioutil.ReadAll(response.Body)
	if response.StatusCode != 201 {
		return errors.New(string(d))
	}

	err = xml.Unmarshal(d, &aitResponse)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

// NewRequest sends the http request to AIT
func (s *ServiceAITImpl) NewRequest(url, apiKey string, method string, data url.Values) (*http.Response, error) {
	req, err := http.NewRequest(method, url, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Add("apiKey", apiKey)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	return s.client.Do(req)
}
