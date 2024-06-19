terraform {
  backend "s3" {
    bucket         = "terraform-backend"
    key            = "asg-ec2-app.state.lock"
    region         = "us-east-1"
    dynamodb_table = "terraform-backend"
  }
}