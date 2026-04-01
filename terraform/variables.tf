variable "api_token" {
  description = "Token simulado para conectarse al API Gateway. Se inyecta desde GitHub Secrets."
  type        = string
  sensitive   = true
  default     = "token-de-prueba-local" # Valor por defecto para que el CI no falle al validar
}
