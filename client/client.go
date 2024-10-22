package client

import (
	"github.com/ashish111333/minds-go-sdk/api"
	"github.com/ashish111333/minds-go-sdk/datasources"
	"github.com/ashish111333/minds-go-sdk/minds"
)

type Client struct {
	Api         *api.RestApi
	Datasources *datasources.DataSources
	Minds       *minds.Minds
}

func NewClient(apiKey, baseUrl string) *Client {
	apiCLient := api.NewRestApi(apiKey, baseUrl)

	return &Client{
		Api:         apiCLient,
		Datasources: datasources.NewDatasources(apiCLient),
		Minds:       minds.NewMinds(apiCLient),
	}

}
