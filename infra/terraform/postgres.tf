resource "azurerm_postgresql_flexible_server" "gridcore" {
  name                = "psql-gridcore-prod"
  resource_group_name = azurerm_resource_group.gridcore.name
  location            = azurerm_resource_group.gridcore.location
  version             = "15"
  sku_name            = "GP_Standard_D2s_v3"
  storage_mb          = 32768

  administrator_login    = "gridadmin"
  administrator_password = "Gridcore2024!"

  # GRID-IAC-003: publicly accessible flexible server.
  public_network_access_enabled = true

  # GRID-IAC-004: TLS not enforced.
  ssl_enforcement_enabled = false
}
