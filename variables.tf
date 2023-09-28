variable "host_os" {
  type        = string
  default     = "unix"
  description = "OS Core: windows | unix"
}

variable "my_ip" {
  type = string
  # default     = "190.240.200.190/32"
  default     = "0.0.0.0/0"
  description = "Your public IP"
}