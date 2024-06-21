output "apigw_name" {
    description = "AṔI Gateway Name"
    value = aws_api_gateway_rest_api.my_api.name
}

output "apigw_url" {
    description = "API Gateway URL"
    value = aws_api_gateway_deployment.apigw_deployment.invoke_url
}