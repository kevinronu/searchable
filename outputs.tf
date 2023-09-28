output "dev-ip" {
  value       = aws_instance.dev_docker.public_ip
  description = "Public IP for your AWS instance"
}
