#!/bin/bash

# Step 1: Build the Docker image using the provided Dockerfile
#echo "Building Docker image..."
#docker build -t my-transactions-app .
echo "Starting MongoDB, RabbitMQ, and the application..."
docker-compose up

# Step 2: Set up the MongoDB database using setup_db.sh script
echo "Setting up MongoDB..."
cd scripts
./setup_db.sh
cd ..

# Step 3: Run the Docker container
#echo "Starting the application..."
#docker run -p 8080:8080 my-transactions-app 

# Step 4: Print the logs to view application output
echo "Viewing application logs..."
docker-compose logs -f

$SHELL