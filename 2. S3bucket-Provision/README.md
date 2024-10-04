Terraform AWS S3 Bucket Configuration
Overview

This repository contains Terraform files for creating and managing an AWS S3 bucket with server-side encryption, versioning, and a bucket policy to allow public read access.
Files
1. main.tf

Defines the main infrastructure:

    Provider: AWS with region specified as ap-southeast-1.
    S3 Bucket:
        Enables versioning.
        Encrypts data using AES256.
        Sets the bucket's ACL to public-read.
    Bucket Policy: Grants public read access for objects in the bucket.

2. output.tf

Outputs relevant bucket details:

    Bucket Name: Outputs the name of the S3 bucket.
    Bucket ARN: Outputs the ARN of the S3 bucket.

3. variables.tf

Defines variables for flexibility:

    Region: Specifies AWS region.
    Bucket Name: Sets the name for the S3 bucket.

Usage

    Initialize Terraform:

    bash

terraform init

Apply the Configuration:

bash

    terraform apply

    Outputs: After applying, you'll see the bucket name and ARN as outputs.

Customization

    Update the region and bucket_name in variables.tf to fit your needs.

Terraform S3 Bucket Deployment with GitHub Actions

This repository automates the deployment of an S3 bucket using Terraform and GitHub Actions.
Workflow Overview

The workflow triggers on every push to the s3bucketprovision branch. It runs the following steps:

    Checkout Code: Retrieves the repository's code.
    Setup Terraform: Installs Terraform version 1.1.7.
    Configure AWS Credentials: Sets AWS credentials from repository secrets.
    Terraform Init: Initializes the Terraform workspace in the terraform/ directory.
    Terraform Plan: Runs terraform plan to preview changes.
    Terraform Apply: Automatically applies the infrastructure changes.

Setup
AWS Credentials

Store your AWS credentials as GitHub secrets:

    AWS_ACCESS_KEY_ID
    AWS_SECRET_ACCESS_KEY
    AWS_REGION

Triggering

Push changes to the s3bucketprovision branch to execute the workflow.
