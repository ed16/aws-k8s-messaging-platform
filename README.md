
Test user-service at localhost:

curl -X POST -H "Content-Type: application/json" -d '{"user_id": "123", "name": "John Doe", "created_at": "2024-01-23"}' http://localhost:8081/create
curl http://localhost:8081/get?name=John%20Doe



It use:
MySQL, PostgreSQL

MongoDB, DynamoDB

Redis 

RabbitMQ, Apache Kafka, SQS


