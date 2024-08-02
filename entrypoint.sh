#!/bin/sh

# Wait for LocalStack to be ready
echo "Waiting for LocalStack to be available..."
until curl -s http://localhost:4566/_localstack/health | grep '"s3": "running"'; do
  sleep 5
done

# Create SQS Queue
echo "Creating SQS queue..."
aws --endpoint-url=http://localhost:4566 sqs create-queue --queue-name calendar-entry

# Start the appropriate service based on an environment variable
case "$SERVICE" in
  "api")
    echo "Starting API service..."
    exec /api
    ;;
  "consumer")
    echo "Starting Consumer service..."
    exec /consumer
    ;;
  "cron")
    echo "Starting Cron service..."
    exec /cron
    ;;
  *)
    echo "No valid service specified, exiting..."
    exit 1
    ;;
esac
