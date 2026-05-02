resource "azurerm_virtual_network" "gridcore" {
  name                = "vnet-gridcore"
  address_space       = ["10.30.0.0/16"]
  location            = azurerm_resource_group.gridcore.location
  resource_group_name = azurerm_resource_group.gridcore.name
}

resource "azurerm_network_security_group" "gridcore" {
  name                = "nsg-gridcore"
  location            = azurerm_resource_group.gridcore.location
  resource_group_name = azurerm_resource_group.gridcore.name

  # GRID-IAC-002: SSH/RDP open to the world.
  security_rule {
    name                       = "allow-mgmt-any"
    priority                   = 100
    direction                  = "Inbound"
    access                     = "Allow"
    protocol                   = "Tcp"
    source_port_range          = "*"
    destination_port_ranges    = ["22", "3389"]
    source_address_prefix      = "*"
    destination_address_prefix = "*"
  }
}
