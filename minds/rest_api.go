package minds

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

var (
	ErrObjectNotFound = errors.New("requested object not found")
	ErrUnknown        = errors.New("unknown error")
	ErrUnauthorized   = errors.New("unauthorized access ")
)

type RestApi struct {
	ApiKey  string
	BaseUrl string
	Client  *http.Client
}

// returns a new instance of RestApi
func NewRestApi(apiKey, baseUrl string) *RestApi {

	if baseUrl == "" {
		baseUrl = "https://mdb.ai"
	}

	return &RestApi{
		ApiKey:  apiKey,
		BaseUrl: baseUrl,
		Client:  &http.Client{Timeout: time.Second * 10},
	}

}

// makes http requests for GET,POST,DELETE,PATCH methods
func (api *RestApi) MakeHttpRequest(httpMethod, url string, RequestData interface{}) (*http.Response, error) {

	// prepare json data if request method is post or patch
	var jsonData []byte
	var err error
	if httpMethod == http.MethodPost || httpMethod == http.MethodPatch {

		jsonData, err = json.Marshal(RequestData)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal json: %w", err)

		}
	}

	//create request
	request, err := http.NewRequest(httpMethod, api.BaseUrl+url, bytes.NewBuffer(jsonData))
	if err != nil {

		return nil, fmt.Errorf("failed to create http request: %w", err)

	}
	//Set Headers

	request.Header.Set("Authorization", "Bearer "+api.ApiKey)
	request.Header.Set("Application-type")
	// make request
	response, err := api.Client.Do(request)
	if err != nil {

		return nil, fmt.Errorf("error in http response : %w", err)
	}
	return response, nil

}

func (api *RestApi) Get() (*http.Response, error) {

}

func (api *RestApi) Delet() (*http.Response, error) {

}

func (api *RestApi) Post() (*http.Response, error) {

}

func (api *RestApi) Patch() (*http.Response, error) {

}
