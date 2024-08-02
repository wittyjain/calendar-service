# Calendar Service

## Overview

The Calendar Service is an application designed to manage calendar entries using MySQL for data storage and SQS for message queuing. This project is set up to run in a Dockerized environment with LocalStack for SQS emulation.

## Features

- **Database Integration**: Utilizes MySQL for managing calendar data.
- **Queue Integration**: Employs SQS (emulated via LocalStack) for queuing and processing messages.
- **Dockerized Environment**: Includes Docker and Docker Compose configuration for seamless deployment and management.

## Prerequisites

- **Docker**: Install Docker from [Docker's official site](https://www.docker.com/get-started).
- **Docker Compose**: Install Docker Compose from [Docker Compose's official site](https://docs.docker.com/compose/install/).

## Getting Started

### Build and Start Containers

Use Docker Compose to build and start the services:

```bash
docker-compose up --build```

This will start the following services:

- LocalStack: For emulating AWS SQS.
- MySQL: For database operations.
- API: Main application service.
- Consumer: Service for processing SQS messages.
- Cron: Service for scheduled tasks.

Configuration
- Configuration is managed via environment variables and YAML files. Update config.yaml with your specific settings:

Sample config.yaml:

```credentials:
  access_key: "test"
  secret_key: "test"
region: "us-east-1"
endpoint: "http://localhost:4566"
queues:
  calendar-entry: "http://localhost:4566/000000000000/calendar-entry"

mysql:
  username: "user"
  password: "password"
  url: "mysql"
  port: "3306"
  database: "calendar"
  charset: "utf8mb4"
  parseTime: "True"
  loc: "Local"
```

Make sure to adjust the credentials and connection details according to your setup.

## Accessing the Services
- API Service: The API service runs on port 8080 by default. You can adjust this in the docker-compose.yml file if needed.
- LocalStack: Accessible on ports 4566 (SQS) and 4571 (other LocalStack services).


## Stopping and Removing Containers
To stop and remove the containers, use:
```docker-compose down```

## Get Active Calendar Entry
```curl --location 'localhost:8080/calendar/active'```

## Process Active Calendar Entries through SQS Queue

```curl --location --request POST 'localhost:8080/process-calendar-entries'```

## Create Active Calendar Entry

```curl --location 'localhost:8080/calendar/active' \
--header 'Content-Type: application/json' \
--data '{
    "StartDate": "2024-08-08T00:00:00Z",
    "StopDate": "2024-08-10T00:00:00Z"
}'```

