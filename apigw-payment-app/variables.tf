variable "project" {
    type = object({
      name = string
      env = string
    })
}

variable "apigateway" {
    type = object({
      path = string
    })
}

variable "lambda" {
  type = object({
    name = string 
  })
}

variable "tags" {
    description = "default tags"
    type = object({
      Name = string
      Environment = string
      Workload = string
      Owner = string
      ProvisionedBy = string
      UpdatedAt = string
      CreatedAt = string 
    })
}