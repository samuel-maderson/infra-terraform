terraform {
  backend "s3" {
    bucket         = "terraform-backend"
    key            = "app-asg-eventbridge.state.lock"
    region         = "us-east-1"
    dynamodb_table = "terraform-backend"
  }
}