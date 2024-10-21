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

func NewClient() *Client {

	return &Client{
		Api:         api.NewRestApi(),
		Datasources: 
		Minds: 
	}

}
