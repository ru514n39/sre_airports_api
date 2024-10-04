# main.tf

# Specify the provider
provider "aws" {
  region = "ap-southeast-1"  # Update this to your desired AWS region
}

# Create the S3 bucket
resource "aws_s3_bucket" "airport_images" {
  bucket = "my-s3-bucket-name"  # Update with your desired unique bucket name

  # Optional: Enable versioning
  versioning {
    enabled = true
  }

  # Optional: Enable server-side encryption
  server_side_encryption_configuration {
    rule {
      apply_server_side_encryption_by_default {
        sse_algorithm = "AES256"
      }
    }
  }

  # Optional: Set access control
  acl = "public-read"  # Make the bucket publicly readable
}

# S3 Bucket Policy to allow public read access
resource "aws_s3_bucket_policy" "bucket_policy" {
  bucket = aws_s3_bucket.airport_images.id

  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Sid       = "PublicReadGetObject",
        Effect    = "Allow",
        Principal = "*",
        Action    = "s3:GetObject",
        Resource  = "${aws_s3_bucket.airport_images.arn}/*"
      }
    ]
  })
}

