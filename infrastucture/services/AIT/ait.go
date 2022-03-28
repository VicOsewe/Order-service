package ait

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/VicOsewe/Order-service/application"
	"github.com/VicOsewe/Order-service/application/dto"
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
	rawRespBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("unable to read wallet response: %w", err)
	}

	aitResponse := dto.SMSResponse{}
	err = json.Unmarshal(rawRespBytes, &aitResponse)
	if err != nil {
		return fmt.Errorf("unable to unmarshal '%s' to JSON: %w", string(rawRespBytes), err)
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
