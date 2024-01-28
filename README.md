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
9. **CI/CD Integration**: Leverage GitHub Actions for continuous integration and deployment processes, ensuring smooth and consistent updates.

## Technologies Used

- **Programming Language**: Go
- **Containerization**: Docker
- **Orchestration**: Kubernetes
- **Databases**: MySQL, PostgreSQL, MongoDB, DynamoDB
- **Message Queues**: RabbitMQ, Apache Kafka, SQS
- **Caching**: Redis
- **CI/CD**: GitHub Actions

This POC is designed to validate the architectural decisions, demonstrating the system's scalability, performance, and reliability, while also providing insights for future optimizations.





## Useful notes:

Test user-service at localhost:

curl -X POST -H "Content-Type: application/json" -d '{"name": "John Doe", "created_at": "2024-01-23"}' http://localhost:8081/create
curl http://localhost:8081/get?name=John%20Doe



Build Docker containers:
cd aws-k8s-messaging-platform
docker build -t load-generator -f services/load-generator/Dockerfile .
docker run -p 8080:8080 load-generator

docker build -t user-service -f services/user-service/Dockerfile .

Docker composer:
docker-compose build
docker-compose up