variable "project" {
  type = object({
    name = string
    region = string
    env = string
  })
}

variable "kms" {
    type = object({
        username = string
        app_user = string
    })
}

variable "tags" {
  type = object({
    Name = string
    Workload = string
    Owner = string
    ProvisionedBy = string
    CreatedAt = string
    UpdatedAt = string
  })
}