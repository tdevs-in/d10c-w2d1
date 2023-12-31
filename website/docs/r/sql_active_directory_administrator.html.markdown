---
subcategory: "Database"
layout: "azurerm"
page_title: "Azure Resource Manager: azurerm_sql_active_directory_administrator"
description: |-
  Manages an Active Directory administrator on a SQL server
---

# azurerm_sql_active_directory_administrator

Allows you to set a user or group as the AD administrator for an Azure SQL server.

-> **Note:** The `azurerm_sql_active_directory_administrator` resource is deprecated in version 3.0 of the AzureRM provider and will be removed in version 4.0. Please use the `azuread_administrator` block of the [`azurerm_mssql_server`](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/mssql_server) resource instead.

## Example Usage

```hcl
data "azurerm_client_config" "current" {}

resource "azurerm_resource_group" "example" {
  name     = "example-resources"
  location = "West Europe"
}

resource "azurerm_sql_server" "example" {
  name                         = "mysqlserver"
  resource_group_name          = azurerm_resource_group.example.name
  location                     = azurerm_resource_group.example.location
  version                      = "12.0"
  administrator_login          = "4dm1n157r470r"
  administrator_login_password = "4-v3ry-53cr37-p455w0rd"
}

resource "azurerm_sql_active_directory_administrator" "example" {
  server_name         = azurerm_sql_server.example.name
  resource_group_name = azurerm_resource_group.example.name
  login               = "sqladmin"
  tenant_id           = data.azurerm_client_config.current.tenant_id
  object_id           = data.azurerm_client_config.current.object_id
}
```

## Argument Reference

The following arguments are supported:

* `server_name` - (Required) The name of the SQL Server on which to set the administrator. Changing this forces a new resource to be created.

* `resource_group_name` - (Required) The name of the resource group for the SQL server. Changing this forces a new resource to be created.

* `login` - (Required) The login name of the principal to set as the server administrator

* `object_id` - (Required) The ID of the principal to set as the server administrator

* `tenant_id` - (Required) The Azure Tenant ID

* `azuread_authentication_only` - (Optional) Specifies whether only AD Users and administrators can be used to login (`true`) or also local database users (`false`).

## Attributes Reference

In addition to the Arguments listed above - the following Attributes are exported:

* `id` - The ID of the SQL Active Directory Administrator.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 30 minutes) Used when creating the SQL Active Directory Administrator.
* `update` - (Defaults to 30 minutes) Used when updating the SQL Active Directory Administrator.
* `read` - (Defaults to 5 minutes) Used when retrieving the SQL Active Directory Administrator.
* `delete` - (Defaults to 30 minutes) Used when deleting the SQL Active Directory Administrator.

## Import

A SQL Active Directory Administrator can be imported using the `resource id`, e.g.

```shell
terraform import azurerm_sql_active_directory_administrator.administrator /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Sql/servers/myserver/administrators/activeDirectory
```
