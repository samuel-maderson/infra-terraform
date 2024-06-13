provider "aws" {
  region = local.region
}

locals {
  region = var.project.region

  tags = var.tags
}

data "aws_caller_identity" "current" {}

resource "aws_kms_alias" "a" {
  name          = "alias/${var.project.env}-${var.project.name}"
  target_key_id = aws_kms_key.my_key.key_id
}

resource "aws_kms_key" "my_key" {
    description             = "An example symmetric encryption KMS key"
    enable_key_rotation     = true
    deletion_window_in_days = 20
    multi_region = true
}

resource "aws_kms_key_policy" "my_key" {
    key_id = aws_kms_key.my_key.id
    policy = jsonencode({
        Version = "2012-10-17"
        Id      = "key-default-1"
        Statement = [
        {
            Sid    = "Enable IAM User Permissions"
            Effect = "Allow"
            Principal = {
            AWS = "arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"
            },
            Action   = "kms:*"
            Resource = "*"
        },
        {
            Sid    = "Allow administration of the key"
            Effect = "Allow"
            Principal = {
            AWS = "arn:aws:iam::${data.aws_caller_identity.current.account_id}:user/${var.kms.username}"
            },
            Action = [
            "kms:ReplicateKey",
            "kms:Create*",
            "kms:Describe*",
            "kms:Enable*",
            "kms:List*",
            "kms:Put*",
            "kms:Update*",
            "kms:Revoke*",
            "kms:Disable*",
            "kms:Get*",
            "kms:Delete*",
            "kms:ScheduleKeyDeletion",
            "kms:CancelKeyDeletion"
            ],
            Resource = "*"
        },
        {
            Sid    = "Allow use of the key"
            Effect = "Allow"
            Principal = {
            AWS = "arn:aws:iam::${data.aws_caller_identity.current.account_id}:user/${var.kms.username}"
            },
            Action = [
            "kms:DescribeKey",
            "kms:Encrypt",
            "kms:Decrypt",
            "kms:ReEncrypt*",
            "kms:GenerateDataKey",
            "kms:GenerateDataKeyWithoutPlaintext"
            ],
            Resource = "*"
        }
        ]
    })
}