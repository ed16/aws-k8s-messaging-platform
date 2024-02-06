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

Project structure:
.
├── .github
│   └── workflows
│       └── ci-cd.yml
├── services
│   ├── load-generator
│   │   ├── cmd
│   │   ├── pkg
│   │   └── Dockerfile
│   ├── user-service
│   │   ├── cmd
│   │   ├── pkg
│   │   └── Dockerfile
│   └── message-service
│       ├── cmd
│       ├── pkg
│       └── Dockerfile
├── web
│   ├── load-generator-ui
│   │   ├── src
│   │   ├── public
│   │   ├── package.json
│   │   └── Dockerfile
│   └── message-user-ui
│       ├── src
│       ├── public
│       ├── package.json
│       └── Dockerfile
├── terraform
│   ├── main.tf
│   ├── variables.tf
│   └── outputs.tf
├── k8s
│   ├── load-generator.yml
│   ├── user-service.yml
│   └── message-service.yml
├── go.mod
├── go.sum
└── README.md


## Learning path:
k8s:
    +implement ingress

    implement local and cloud volume. Use persistentVolumeClaim
    deploy stateful services for DB: manually and using operators

    Use Helm charts when the landscape become too complex

## Useful notes:

Test user-service at localhost:

curl -X POST -H "Content-Type: application/json" -d '{"name": "John Doe", "created_at": "2024-01-23"}' http://localhost:8081/create
curl http://localhost:8081/get?id=1


Build Docker containers:

docker build -t ed16/aws-k8s-messaging-platform:user-service-v0.3 -f services/user-service/Dockerfile .
docker build -t ed16/aws-k8s-messaging-platform:load-generator-v0.3 -f services/load-generator/Dockerfile .
docker push ed16/aws-k8s-messaging-platform:user-service-v0.3
docker push ed16/aws-k8s-messaging-platform:load-generator-v0.3

minikube ssh -p minikube

### Prometheus
minikube service prometheus-service -n prometheus

kubectl rollout restart deployment prometheus-deployment -n prometheus

Prometheus server URL for Grafana:
http://prometheus-service.prometheus.svc.cluster.local:9090

## TODO

1. +Write 2 go services
2. +Wrap into docker containers
3. +Deploy into minikube
4. +Test in minikube
5. Fix minikube tunnel


## Deploy from zero
minikube start
find k8s -name '*.yaml' | xargs -I {} kubectl apply -f {}
kubectl get all
kubectl get all -n prometheus
kubectl get all -n grafana