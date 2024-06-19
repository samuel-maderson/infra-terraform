# __generated__ by Terraform
# Please review these resources and move them into your main configuration files.

# __generated__ by Terraform from "test-automation/Idbb0af53a-5f06-4dd2-bce1-9b378eaa45ad"
resource "aws_cloudwatch_event_target" "default" {
  arn            = var.eventbrid_rule.arn_target
  event_bus_name = "default"
  force_destroy  = false
  input          = null
  input_path     = null
  role_arn       = null
  rule           = var.project.name
  target_id      = var.eventbrid_rule.target_id
}
