terraform { required_version = ">= 1.6" }
resource "aws_vpc" "this" { cidr_block = var.cidr }
