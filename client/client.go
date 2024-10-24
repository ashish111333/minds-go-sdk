package client

import (
	"github.com/ashish111333/minds-go-sdk/api"
	"github.com/ashish111333/minds-go-sdk/datasources"
	"github.com/ashish111333/minds-go-sdk/minds"
)

type Client struct {
	Api         *api.RestApi
	Datasources *datasources.DataSources
	Minds       *minds.Mind
}
type ApiConfig struct {
	ApiKey  string
	BaseUrl string
}

func NewClient(ac ApiConfig) *Client {

	apiCLient := api.NewRestApi(ac.ApiKey, ac.BaseUrl)

	return &Client{
		Api:         apiCLient,
		Datasources: datasources.NewDatasources(apiCLient),
		Minds:       minds.NewMindsClient(apiCLient),
	}

}
