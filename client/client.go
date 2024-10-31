package client

import (
	"fmt"

	"github.com/ashish111333/minds-go-sdk/api"
	"github.com/ashish111333/minds-go-sdk/datasources"
	"github.com/ashish111333/minds-go-sdk/minds"
)

type Client struct {
	Api         *api.RestApi
	Datasources *datasources.DataSources
	Minds       *minds.Mind
}

func NewClient(apiKey string, baseUrl ...string) (*Client, error) {

	var base_url string
	if len(baseUrl) == 0 {
		base_url = ""
	} else {
		base_url = baseUrl[0]
	}
	if apiKey == "" {
		return nil, fmt.Errorf("api key cant be empty")
	}
	apiCLient := api.NewRestApi(apiKey, base_url)

	return &Client{
		Api:         apiCLient,
		Datasources: datasources.NewDatasourcesClient(apiCLient),
		Minds:       minds.NewMindsClient(apiCLient),
	}, nil

}
