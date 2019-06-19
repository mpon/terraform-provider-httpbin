package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Provider is the root of terraform provider plugin
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{},
	}
}
