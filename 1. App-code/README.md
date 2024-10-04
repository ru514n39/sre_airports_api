App Deployment Flows

This repository contains an automated workflow for building and deploying a Go-based application using Docker.
Workflow Overview

The GitHub Actions workflow is triggered by any push to the main branch. It includes the following steps:

    Checkout Code: Clones the repository.
    Setup Go: Installs Go version 1.23.1.
    Install Dependencies: Downloads required Go packages for AWS SDK.
    Build & Push Docker Image:
        Builds a Docker image using app.dockerfile.
        Tags the image using the GitHub run ID.
        Pushes the image to Docker Hub.

Setup
Docker Hub Credentials

Store your Docker Hub username and password in GitHub Secrets:

    DOCKER_USER
    DOCKER_PASSWORD

Triggering

Push changes to the main branch to execute the workflow.

Dockerfile Overview

This repository contains a multi-stage Dockerfile for building and running a Go application that integrates with AWS services such as S3.
Stages

    Build Stage (Golang 1.23.1)
        Sets up the Go environment.
        Downloads dependencies like aws-sdk-go.
        Compiles the Go binary.

    Final Stage (Ubuntu)
        Copies the compiled binary from the build stage.
        Runs the application in a minimal Ubuntu image.
        Exposes port 8080.



Go Application with S3 Image Upload and Airport API

This project implements a simple Go-based API that serves airport data, includes health checks, and supports image uploads to AWS S3 for airport updates.
Features

    Airport API: Provides data on airports in Bangladesh.
    S3 Image Upload: Allows users to upload airport images to an S3 bucket.
    Endpoints:
        /: Home page
        /airports: Get basic airport information
        /airports_v2: Get airport info with runway lengths
        /healthcheck: Health check endpoint
        /update_airport_image: Upload an image for an airport

Getting Started
Prerequisites

    Go 1.23.1
    AWS Account with S3 access

Run the Application

    Clone the repository:

    bash

git clone <repo_url>

Set up your AWS credentials.
Build and run the application:

bash

    go run main.go

Endpoints

    Airports
        URL: /airports
        Method: GET
        Returns a list of airports with their respective data.

    Airports V2
        URL: /airports_v2
        Method: GET
        Returns a list of airports with runway lengths.

    Update Airport Image
        URL: /update_airport_image
        Method: POST
        Uploads a new image for an airport and updates the S3 bucket.

Image Upload to S3

The application supports uploading airport images to an AWS S3 bucket via the /update_airport_image endpoint. Upon successful upload, the S3 URL is returned and stored with the airport data.
Logging

Logs are stored in server.log, capturing server startup times, requests, and responses.
Deployment

    The application runs on port 8080.
    Deploy the application using Docker or directly on a server that supports Go binaries.

Example

Here is an example of how to use the /update_airport_image endpoint to upload an image for a specific airport:

bash

curl -X POST -F "name=Hazrat Shahjalal International Airport" -F "image=@path/to/image.jpg" http://localhost:8080/update_airport_image

The response will include the updated image URL.
