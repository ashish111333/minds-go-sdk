package minds

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	ErrObjectNotFound = errors.New("requested object not found")
	ErrUnknown        = errors.New("unknown error")
	ErrUnauthorized   = errors.New("unauthorized access ")
	ErrForbidden      = errors.New("access forbidden")
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
	headers := map[string]string{
		"Authorization": "Bearer " + api.ApiKey,
		"Content-Type":  "application/json",
	}
	api.setHeaders(request, headers)

	// make request
	response, err := api.Client.Do(request)
	if err != nil {

		return nil, fmt.Errorf("error in http response : %w", err)
	}
	return response, nil

}

// helper function for http status code errors
func (api *RestApi) handleErrorForStatus(response *http.Response) error {
	defer response.Body.Close()
	switch response.StatusCode {
	case 404:
		return ErrObjectNotFound
	case 403:
		return ErrForbidden
	case 401:
		return ErrUnauthorized

	default:
		if response.StatusCode >= 400 && response.StatusCode < 600 {
			return ErrUnknown
		}

	}
	return nil

}

// helper function to set headers for a given request
func (api *RestApi) setHeaders(request *http.Request, headers map[string]string) {
	for key, value := range headers {
		request.Header.Set(key, value)
	}

}

// GET Method
func (api *RestApi) Get(url string, Requestdata interface{}) (*http.Response, error) {
	log.Printf("making get request Url: %s Data: %v \n", url, Requestdata)
	response, err := api.MakeHttpRequest(http.MethodGet, url, Requestdata)
	if err != nil {

		log.Printf("http get request failed : %v \n", err)
		return nil, err
	}
	err = api.handleErrorForStatus(response)
	if err != nil {
		log.Printf("error in http response:  %v \n", err)
		return nil, err
	}
	return response, nil
}

// DELETE method
func (api *RestApi) Delete(url string, Requestdata interface{}) (*http.Response, error) {
	log.Printf("making Delete request Url: %s Data: %v \n", url, Requestdata)
	response, err := api.MakeHttpRequest(http.MethodDelete, url, Requestdata)
	if err != nil {
		log.Printf("http Delete request failed :%v  ", err)
		return nil, err

	}
	err = api.handleErrorForStatus(response)
	if err != nil {
		log.Printf("error in http response %v \n", err)
		return nil, err

	}

}

// POST method
func (api *RestApi) Post(url string, Requestdata interface{}) (*http.Response, error) {

	log.Printf("making post request to %s , data : %v \n", url, Requestdata)
	response, err := api.MakeHttpRequest(http.MethodPost, url, Requestdata)
	if err != nil {

		log.Printf("failed to make Http Post request : %v \n", err)
		return nil, err

	}
	err = api.handleErrorForStatus(response)
	if err != nil {

		log.Printf("error in http response: %v \n", err)
		return nil, err
	}
	return response, nil

}

// PATCH method
func (api *RestApi) Patch(url string, Requestdata interface{}) (*http.Response, error) {
	response, err := api.MakeHttpRequest(http.MethodPatch, url, Requestdata)
	if err != nil {

		log.Printf("failed to make Http Patch request: %v \n", err)
		return nil, err
	}
	err = api.handleErrorForStatus(response)
	if err != nil {
		log.Printf("error in http response: %v \n", err)
		return nil, err
	}
	return response, nil

}
