data "aws_vpc" "default_vpc_data"{
    default = true

}


resource "aws_security_group" "web_sg" {
    name        = "web_sg"
    vpc_id      = data.aws_vpc.default_vpc_data.id

    tags = {
      "Name"    = "web_sg"
    }
  
}

resource "aws_security_group_rule" "allow_ssh" {
    type = "ingress"
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
    security_group_id = aws_security_group.web_sg.id
}

resource "aws_security_group_rule" "allow_out" {
    type = "egress"
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
    security_group_id = aws_security_group.web_sg.id
}
resource "aws_security_group_rule" "allow_http" {
    type = "ingress"
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
    security_group_id = aws_security_group.web_sg.id
}

resource "aws_security_group" "db_sg" {
    name        = "db_sg"
    vpc_id      = data.aws_vpc.default_vpc_data.id

    tags = {
      "Name"    = "db_sg"
    }
  
}

resource "aws_security_group_rule" "allow_db" {
    type = "ingress"
    from_port   = 3306
    to_port     = 3306
    protocol    = "tcp"
    source_security_group_id = aws_security_group.web_sg.id
    security_group_id = aws_security_group.db_sg.id
}

resource "aws_security_group_rule" "allow_db_out" {
    type = "egress"
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
    security_group_id = aws_security_group.db_sg.id
}