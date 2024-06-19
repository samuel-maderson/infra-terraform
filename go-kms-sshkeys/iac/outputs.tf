output "key_id" {
  value = aws_kms_key.my_key.id
}

output "key_arn" {
  value = aws_kms_key.my_key.arn
}

output "secret_name" {
  value = aws_secretsmanager_secret.secret.name
}