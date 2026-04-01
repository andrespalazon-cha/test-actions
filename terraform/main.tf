terraform {
  required_version = ">= 1.0.0"

  required_providers {
    random = {
      source  = "hashicorp/random"
      version = "~> 3.6.0"
    }
  }
}

# Un recurso de prueba que genera un nombre aleatorio (ej. "upward-coyote")
resource "random_pet" "gateway_name" {
  length    = 2
  separator = "-"
}
