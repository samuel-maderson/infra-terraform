
# Create a secret in AWS Secrets Manager
resource "aws_secretsmanager_secret" "secret" {
  name        = "${var.project.env}-${var.project.name}"
  description = "my rds secret"
}