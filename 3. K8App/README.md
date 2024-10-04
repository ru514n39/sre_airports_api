Airport Application Deployment with ArgoCD and Kubernetes
Overview

This repository contains the necessary configuration files to deploy and manage the Airport Application using ArgoCD, Kubernetes, and Helm. The deployment process includes automated updates using the ArgoCD Image Updater and monitors the application with Kubernetes probes.
Files Overview
1. ArgoCD Application Manifest (Application.yaml)

Defines the ArgoCD application:

    Name: airport-application
    Namespace: argocd
    Source: GitHub repository (goz/airport-app).
    Sync Policy: Automatic sync, pruning, and self-healing enabled.
    Image Updater: Automatically tracks the image tag using annotations.

2. Kubernetes Deployment Manifest (Deployment.yaml)

Specifies the deployment of the Airport Application:

    Replicas: 3 instances.
    Container: kylo39/doom image.
    Resource Requests & Limits: Configured for efficient resource use.
    Probes: Readiness and liveness probes configured for health checks.

3. Kubernetes Service Manifest (Service.yaml)

Defines the service to expose the Airport Application:

    Type: LoadBalancer (accessible externally).
    Ports: External port 80 mapped to internal container port 8080.

4. Helm Values (values.yaml)

Holds the values for image name and tag for dynamic updates during deployments.
Setup Instructions
Prerequisites

    ArgoCD installed on your Kubernetes cluster.
    Proper Git credentials set up for ArgoCD.
    Helm installed.

Deploying the Application

    Clone the Repository:

    bash

git clone git@github.com:ru514n39/goz.git
cd goz/airport-app

Apply ArgoCD Application:

bash

kubectl apply -f Application.yaml

Monitor Deployment: Open ArgoCD UI and watch for the sync status of airport-application. Ensure that the application is synced and healthy.

Access the Application: Once the service is up, access the application using the external IP provided by the LoadBalancer:

bash

    kubectl get svc airport-app-service

Automatic Image Updates

    The application automatically updates to the latest version of the image kylo39/doom using the ArgoCD Image Updater annotations.
    The new image is pulled based on the latest tag pushed to the repository.
