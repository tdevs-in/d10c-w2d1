// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apimanagement_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azurerm/internal/clients"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/apimanagement/parse"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/utils"
)

type ApiManagementApiTagDescriptionResource struct{}

func TestAccApiManagementApiTagDescription_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_api_management_api_tag_description", "test")
	r := ApiManagementApiTagDescriptionResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApiManagementApiTagDescription_requiresImport(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_api_management_api_tag_description", "test")
	r := ApiManagementApiTagDescriptionResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.RequiresImportErrorStep(r.requiresImport),
	})
}

func TestAccApiManagementApiTagDescription_update(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_api_management_api_tag_description", "test")
	r := ApiManagementApiTagDescriptionResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.update(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func (ApiManagementApiTagDescriptionResource) Exists(ctx context.Context, clients *clients.Client, state *pluginsdk.InstanceState) (*bool, error) {
	id, err := parse.ApiTagDescriptionsID(state.ID)
	if err != nil {
		return nil, err
	}

	resp, err := clients.ApiManagement.ApiTagDescriptionClient.Get(ctx, id.ResourceGroup, id.ServiceName, id.ApiName, id.TagDescriptionName)
	if err != nil {
		return nil, fmt.Errorf("reading %q: %+v", id, err)
	}

	return utils.Bool(resp.ID != nil), nil
}

func (r ApiManagementApiTagDescriptionResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s
resource "azurerm_api_management_tag" "test" {
  api_management_id = azurerm_api_management.test.id
  name              = "acctest-Tag-%d"
}

resource "azurerm_api_management_api_tag" "test" {
  api_id = azurerm_api_management_api.test.id
  name   = "acctest-Tag-%d"
}

resource "azurerm_api_management_api_tag_description" "test" {
  api_tag_id                         = azurerm_api_management_api_tag.test.id
  description                        = "tag description"
  external_documentation_url         = "https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs"
  external_documentation_description = "external tag description"
}
`, ApiManagementApiResource{}.basic(data), data.RandomInteger, data.RandomInteger)
}

func (r ApiManagementApiTagDescriptionResource) requiresImport(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s
resource "azurerm_api_management_api_tag_description" "import" {
  api_tag_id                         = azurerm_api_management_api_tag_description.test.api_tag_id
  description                        = azurerm_api_management_api_tag_description.test.description
  external_documentation_url         = azurerm_api_management_api_tag_description.test.external_documentation_url
  external_documentation_description = azurerm_api_management_api_tag_description.test.external_documentation_description
}
`, r.basic(data))
}

func (r ApiManagementApiTagDescriptionResource) update(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s
resource "azurerm_api_management_tag" "test" {
  api_management_id = azurerm_api_management.test.id
  name              = "acctest-Tag-%d"
}

resource "azurerm_api_management_api_tag" "test" {
  api_id = azurerm_api_management_api.test.id
  name   = "acctest-Tag-%d"
}

resource "azurerm_api_management_api_tag_description" "test" {
  api_tag_id                         = azurerm_api_management_api_tag.test.id
  description                        = "tag description update"
  external_documentation_url         = "https://registry.terraform.io/providers/hashicorp/azurerm"
  external_documentation_description = "external tag description update"
}
`, ApiManagementApiResource{}.basic(data), data.RandomInteger, data.RandomInteger)
}
