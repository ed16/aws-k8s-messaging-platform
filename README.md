[![build](https://github.com/ed16/aws-k8s-messaging-platform/actions/workflows/ci-cd.yml/badge.svg)](https://github.com/ed16/aws-k8s-messaging-platform/actions/workflows/ci-cd.yml)
[![Coverage Status](https://coveralls.io/repos/github/ed16/aws-k8s-messaging-platform/badge.svg)](https://coveralls.io/github/ed16/aws-k8s-messaging-platform)

# Microservices-Based Proof of Concept (POC) Project

This project showcases a microservices architecture, primarily developed using Go and efficiently containerized with Docker. It includes a collection of services, each tailored to fulfill specific roles within the ecosystem.

## Services Overview

### User Service

- **`POST user-service/create`**: Registers a new user.
  - **Payload Example**: `{"name": "John Doe", "created_at": "2024-01-23"}`
  - **Response**: No content.
  
- **`GET user-service/get`**: Retrieves user details by ID.
  - **Parameters**: `id`
  - **Response Example**: `{"user_id": "123", "name": "John Doe", "created_at": "2024-01-23"}`

### Load Generator Service

- **`GET load-generator/start`**: Triggers payload generation for user creation.
  - **Parameters**: None
  - **Response**: No content.

## Project Objectives

1. **Kubernetes Autoscaling Validation**: Assess the auto-scaling feature of Kubernetes under various load conditions.
2. **Database Management and Analysis**: Utilize different databases (MySQL, PostgreSQL, MongoDB, and DynamoDB) to measure and compare throughput, performance, scalability, and reliability.
3. **Asynchronous Communication**: Explore the use and optimization of message queues (RabbitMQ, Apache Kafka, SQS) in handling asynchronous requests.
4. **Caching and Performance Enhancement**: Examine the impact of Redis caching on performance enhancement, especially for slow requests and complex data structures.
5. **Monitoring**: Implement comprehensive monitoring systems for real-time visibility into system health and performance.
6. **Logging**: Establish robust logging mechanisms for operational transparency and effective troubleshooting.
7. **Security**: Implement stringent security protocols to protect the system and data.
8. **Kubernetes Deployment**: Utilize Kubernetes for the efficient, scalable deployment of services.
9. **Infrastructure as Code (IaC)**: Leverage Terraform to define and provision the cloud infrastructure in a consistent and repeatable manner.
10. **CI/CD Integration**: Leverage GitHub Actions for continuous integration and deployment processes, ensuring smooth and consistent updates.

## Technologies Used

- **Programming Language**: Go
- **Containerization**: Docker
- **Orchestration**: Kubernetes
- **Infrastructure as Code (IaC)**: Terraform
- **Databases**: MySQL, PostgreSQL, MongoDB, DynamoDB
- **Message Queues**: RabbitMQ, Apache Kafka, SQS
- **Caching**: Redis
- **CI/CD**: GitHub Actions

This POC is designed to validate the architectural decisions, demonstrating the system's scalability, performance, and reliability, while also providing insights for future optimizations.

## Learning path:
k8s:
  Use Persistent Volumes for Grafana config
  deploy stateful services for DB: manually and using operators
  Use Helm charts when the landscape become too complex

## Useful notes:

Test user-service at localhost:
curl -X POST -H "Content-Type: application/json" -d '{"name": "John Doe", "created_at": "2024-01-23"}' http://localhost:8081/create
curl http://localhost:8081/get?id=1

Execute linter checks:
golangci-lint run

Build Docker containers, push to the repository, deploy to k8s:

docker build -t ed16/aws-k8s-messaging-platform:user-service-latest -f services/user-service/Dockerfile .
docker build -t ed16/aws-k8s-messaging-platform:load-generator-latest -f services/load-generator/Dockerfile .
docker push ed16/aws-k8s-messaging-platform:user-service-latest
docker push ed16/aws-k8s-messaging-platform:load-generator-latest

kubectl rollout restart deployment user-service 
kubectl rollout restart deployment load-generator

minikube ssh -p minikube

find . -type f -name "*.go" -exec sh -c 'echo "File: {}"; echo "----------------"; cat "{}"; echo "\n"' \;
bombardier -c 64 -n 100000000 -m POST -t 5s -f ./bombardier/payload.json -H "Content-Type: application/json" http://127.0.0.1:8080/create

Vertical Pod Autoscaler
https://github.com/kubernetes/autoscaler/tree/master/vertical-pod-autoscaler#installation

Metrics Server (required for Horisontal Pods Autoscaler)
https://github.com/kubernetes-sigs/metrics-server?tab=readme-ov-file#installation

### Prometheus
minikube service prometheus-service -n prometheus

kubectl rollout restart deployment prometheus-deployment -n prometheus

Prometheus server URL for Grafana:
http://prometheus-service.prometheus.svc.cluster.local:9090

## TODO

1. +Write 2 go services
2. +Wrap into docker containers
3. +Deploy into minikube
4. +Deploy Prometheus and Grafana
5. +Collect custom metrics - user creation rate
6. +Display custom metrics in Grafana
7. +Implement number of processes control for load-generator
8. +Each new commit to the main branch (github actions)
    1.  Run unit tests
    2.  Build containers
    3.  Run integrational tests
9.  Horizontal autoscaling of the user-service
10. Vertical autoscaling of the load-generator
11. Store users in PostgresDB
12. Store users in MongoDB


## Deploy from zero
brew install docker
brew install colima
brew install kubernetes-cli
brew install minikube
colima start --cpu 8 --memory 16
minikube start --driver=docker --cpus=4 --memory=8192 --disk-size=30g
minikube addons enable ingress
kubectl create secret generic grafana-api-token --from-literal=apiToken='<Your Grafana.com API Token>' -n prometheus
kubectl create secret generic postgres-secret \
  --from-literal=POSTGRES_USER=user \
  --from-literal=POSTGRES_PASSWORD=password \
  --from-literal=POSTGRES_DB=mydatabase

kubectl create secret generic mongodb-secret --from-literal=mongo-root-username='username' --from-literal=mongo-root-password='password'
  
find k8s -name '*.yaml' | xargs -I {} kubectl apply -f {}
helm repo add open-telemetry https://open-telemetry.github.io/opentelemetry-helm-charts
helm repo update
helm install opentelemetry-collector open-telemetry/opentelemetry-collector --set mode=daemonset

kubectl get all
kubectl get all -n prometheus
kubectl get all -n grafana




curl -I http://minikube.local/load-generator/SetCreateUsersConnections?c=0