variable "project" {
  description = "main values"
  type = object({
    name = string
    region = string 
    env = string
  })
}

variable "vpc" {
    description = "my vpc name"
    type = object({
      name = string
      private_subnets_pattern = string
    })
}

variable "asg" {
    description = "my asg name"
    type = object({
      min_size = number
      max_size = number
      desired_capacity = number
      instance_type = string
      ami_name = string
      ami_owner = string
      target_group_arn = string
      ec2_ni_private_subnet = string
      kms_key_id = string
      health_check_grace_period = number
    })
}

variable "tags" {
  description = "my tags"
  type = object({
    Name = string
    Workload = string
    Owner = string
    ProvisionedBy = string
    CreatedAt = string
    UpdatedAt = string 
  })
}