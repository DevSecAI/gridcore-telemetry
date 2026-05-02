terraform { required_version = ">= 1.5" }
provider "azurerm" { features {} }

resource "azurerm_resource_group" "gridcore" {
  name     = "rg-gridcore-prod"
  location = "uksouth"
}

# GRID-IAC-001: storage account allows public network access and anonymous blobs.
resource "azurerm_storage_account" "telemetry" {
  name                     = "gridcoretelemetryprod"
  resource_group_name      = azurerm_resource_group.gridcore.name
  location                 = azurerm_resource_group.gridcore.location
  account_tier             = "Standard"
  account_replication_type = "LRS"

  public_network_access_enabled   = true
  allow_nested_items_to_be_public = true
  min_tls_version                 = "TLS1_0"
}
