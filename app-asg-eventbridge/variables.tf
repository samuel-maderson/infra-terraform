variable "project" {
    description = "project global variables"
    type = object({
      name = string
      env = string
    })
}

variable "eventbrid_rule" {
    description = "eventbrid rule"
    type = object({
      arn_target = string
      asg_name = string
      target_id = string
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