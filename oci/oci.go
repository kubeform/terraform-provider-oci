package oci

import (
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/terraform-providers/terraform-provider-oci/internal/provider"
)

func Provider() terraform.ResourceProvider {
	return provider.Provider()
}
