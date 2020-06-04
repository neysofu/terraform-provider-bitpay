package bitpay

import (
	"github.com/bitpay/bitpay-go/client"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var version string

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema:         map[string]*schema.Schema{},
		ConfigureFunc:  configureProvider,
		ResourcesMap:   map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{},
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	return providerMeta{nil}, nil
}

type providerMeta struct {
	client *client.Client
}
