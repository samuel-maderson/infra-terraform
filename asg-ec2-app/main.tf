provider "aws" {
  region = local.region
}

data "aws_availability_zones" "available" {}
data "aws_ami" "my_ami" {
  most_recent = true
  owners      = [var.asg.ami_owner]

  filter {
    name   = "name"
    values = [var.asg.ami_name]
  }
}

data "aws_vpc" "selected" {
    filter {
        name = "tag:Name"
        values = [var.vpc.name]
    }
}

data "aws_subnets" "main_ones" {
    filter {
        name = "vpc-id"
        values = [data.aws_vpc.selected.id]
    }

    tags = {
        Name = var.vpc.private_subnets_pattern
    }
}

data "aws_lb_target_group" "tg_test" {
    arn  = var.asg.target_group_arn
}

locals {
  name   = "${var.project.env}-${var.project.name}"
  region = var.project.region

  #vpc_cidr = var.vpc.cidr
  azs      = slice(data.aws_availability_zones.available.names, 0, 3)

  tags = var.tags
}

module "autoscaling" {
  source  = "terraform-aws-modules/autoscaling/aws"
  version = "7.4.1"
  # insert the 1 required variable here

  # Autoscaling group
  name            = "${var.project.env}-${var.project.name}"
  use_name_prefix = false
  instance_name   = local.name

  ignore_desired_capacity_changes = true

  min_size                  = var.asg.min_size
  max_size                  = var.asg.max_size
  desired_capacity          = var.asg.desired_capacity
  wait_for_capacity_timeout = 0
  # default_instance_warmup   = 300
  health_check_type         = "ELB"
  health_check_grace_period = var.asg.health_check_grace_period
  vpc_zone_identifier       = data.aws_subnets.main_ones.ids
  service_linked_role_arn   = aws_iam_service_linked_role.autoscaling.arn

  # Traffic source attachment
  create_traffic_source_attachment = true
  traffic_source_identifier        = data.aws_lb_target_group.tg_test.arn
  traffic_source_type              = "elbv2"

  initial_lifecycle_hooks = [
    {
      name                 = "delay"
      default_result       = "CONTINUE"
      heartbeat_timeout    = 120
      lifecycle_transition = "autoscaling:EC2_INSTANCE_LAUNCHING"
      # This could be a rendered data resource
      #notification_metadata = jsonencode({ "hello" = "world" })
    }
  ]

  # Launch template
  launch_template_name        = "${var.project.env}-${var.project.name}"
  launch_template_description = "My launch template"
  update_default_version      = true

  image_id          = data.aws_ami.my_ami.id
  instance_type     = var.asg.instance_type
#   user_data         = filebase64("scripts/user_data.sh")
  ebs_optimized     = true
  enable_monitoring = true

  create_iam_instance_profile = true
  iam_role_name               = "${var.project.env}-${var.project.name}"
  iam_role_path               = "/ec2/"
  iam_role_description        = "My IAM Role description"
  iam_role_tags = {
    CustomIamRole = "Yes"
  }
  iam_role_policies = {
    AmazonSSMManagedInstanceCore = "arn:aws:iam::aws:policy/AmazonSSMManagedInstanceCore"
    S3FullAccess = "arn:aws:iam::aws:policy/AmazonS3FullAccess"
    CloudWatchAgentServerPolicy = "arn:aws:iam::aws:policy/CloudWatchAgentServerPolicy"
  }

  # Security group is set on the ENIs below
  security_groups          = [module.asg_sg.security_group_id]

  block_device_mappings = [
    {
      device_name = "/dev/sda1"
      no_device   = 1
      ebs = {
        delete_on_termination = true
        encrypted             = true
        volume_size           = 50
        volume_type           = "gp3"
        kms_key_id  = var.asg.kms_key_id
      }
    }
  ]

  network_interfaces = [
    {
      subnet_id = var.asg.ec2_ni_private_subnet
      delete_on_termination = true
      description           = "eth0"
      device_index          = 0
      security_groups       = [module.asg_sg.security_group_id]
    }
  ]

  
  placement = {
    availability_zone = "${local.azs[2]}"
  }

  
  # Target scaling policy schedule based on average CPU load
  scaling_policies = {
    avg-cpu-policy-greater-than-25 = {
      policy_type               = "TargetTrackingScaling"
      estimated_instance_warmup = 60
      target_tracking_configuration = {
        predefined_metric_specification = {
          predefined_metric_type = "ASGAverageCPUUtilization"
        }
        target_value = 25.0
      }
    }
  }
}


module "asg_sg" {
  source  = "terraform-aws-modules/security-group/aws"
  version = "~> 5.0"

  name        = local.name
  description = "A security group"
  vpc_id      = data.aws_vpc.selected.id

  ingress_with_cidr_blocks = [
    {
      from_port        = 80
      to_port          = 80
      protocol         = "tcp"
      description      = "HTTP"
      cidr_blocks = "0.0.0.0/0"
    },
    {
      from_port        = 22
      to_port          = 22
      protocol         = "tcp"
      description      = "SSH"
      cidr_blocks = "0.0.0.0/0"
    },
    {
      from_port        = -1
      to_port          = -1
      protocol         = "icmp"
      description      = "ICMP"
      cidr_blocks = "0.0.0.0/0"
    }
  ]

  egress_rules = ["all-all"]

  tags = local.tags
}

resource "aws_iam_service_linked_role" "autoscaling" {
  aws_service_name = "autoscaling.amazonaws.com"
  description      = "A service linked role for autoscaling"
  custom_suffix    = local.name

  # Sometimes good sleep is required to have some IAM resources created before they can be used
  provisioner "local-exec" {
    command = "sleep 10"
  }
}