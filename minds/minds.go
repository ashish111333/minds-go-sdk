package minds

import (
	"github.com/ashish111333/minds-go-sdk/datasources"
)

type Mind struct {
	ModeName    string
	provider    string
	CreatedAt   string
	UpdatedAt   string
	datasources *datasources.DataSource
}
