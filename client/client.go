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

func NewClient(apiKey, baseUrl string) (*Client, error) {

	if apiKey == "" {
		return nil, fmt.Errorf("Api key cant be empty")
	}
	apiCLient := api.NewRestApi(apiKey, baseUrl)

	return &Client{
		Api:         apiCLient,
		Datasources: datasources.NewDatasources(apiCLient),
		Minds:       minds.NewMindsClient(apiCLient),
	}, nil

}
