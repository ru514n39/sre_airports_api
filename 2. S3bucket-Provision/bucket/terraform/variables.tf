# variables.tf

variable "region" {
  default = "ap-southeast-1"
}

variable "bucket_name" {
  description = "The name of the S3 bucket"
  default     = "my-s3-bucket-name"  # Replace with your bucket name
}
