resource "aws_cloudwatch_event_rule" "console" {
  description    = null
  event_bus_name = "default"
  event_pattern = jsonencode({
    detail = {
      AutoScalingGroupName = [var.eventbrid_rule.asg_name]
    }
    detail-type = ["EC2 Instance-launch Lifecycle Action"]
    source      = ["aws.autoscaling"]
  })
  force_destroy       = false
  name                = var.project.name
  name_prefix         = null
  role_arn            = null
  schedule_expression = null
  state               = "ENABLED"
  tags                = {}
  tags_all            = {}
}