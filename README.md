# Go Assessment Project
Welcome to the Go Assessment Project! This repository showcases a Go-based application built with a microservices architecture. Follow the instructions below to set up and run the project effortlessly.

## 🚀 Prerequisites
Ensure you have the following installed:

1. Docker
2. Docker Compose

## ⚙️ Setup Instructions
1. Clone the Repository
2. Navigate to the Project Root
3. In the root directory, create a .env file to configure your environment variables. Here's a default setup you can copy and paste:
```
DB_HOST=localhost
POSTGRES_USER=postgres
POSTGRES_PASSWORD=password
POSTGRES_DB=ride_hailing
PORT=5432
```
4. Launch the application with Docker Compose ``docker-compose up``

That's it! 🚀 All services will be set up automatically, and the database will be created with the required schema pre-initialized.

## 📡 Using the Services
Once the services are running, you can start sending gRPC requests to the following default ports on your host machine:

1. User Service: localhost:8090
2. Ride Service: localhost:8088
3. Booking Service: localhost:8092

## 💡 Bonus Feature
I’ve added a microservice for the Ride Service as well.

Have fun, and happy coding! 🚀🎉