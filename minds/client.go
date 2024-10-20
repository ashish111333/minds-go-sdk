package minds

import (
	"github.com/ashish111333/minds-go-sdk/minds/datasources"
)

type Client struct {
	Api         *minds.RestApi
	Datasources datasources.DataSources
}

func NewClient() {

}
