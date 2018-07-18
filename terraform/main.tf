resource "azurerm_resource_group" "clippy-rg" {
  name = "clippy"
  location = "west us"
}

resource "azurerm_container_group" "clippy-container" {
  name                = "clippybrain"
  location            = "${azurerm_resource_group.clippy-rg.location}"
  resource_group_name = "${azurerm_resource_group.clippy-rg.name}"
  ip_address_type     = "public"
  os_type             = "linux"

# a snapshot of the project from early March is used as the image, update if needed
  container {
    name   = "clippybrain"
    image  = "jluk/clippy-bot"
    cpu    = "1"
    memory = "0.5"
    port = "80"

    environment_variables {
      API_KEY="<YOUR-SLACK-API-KEY>"
      BOT_USER="<USERNAME>"
      KEY_WORD="<KEYWORD>"
    }
  }

}