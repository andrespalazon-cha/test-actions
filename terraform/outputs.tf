output "gateway_assigned_name" {
  description = "El nombre aleatorio generado para nuestro Gateway simulado."
  value       = random_pet.gateway_name.id
}
