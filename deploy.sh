#!/bin/bash

set -e

# Change to the project directory
cd /home/ifkash/boo/

# Pull the latest changes
git pull origin main

# Stop and remove previous containers
docker compose down -v

# Build and start new containers
docker compose up --build -d

echo "Deployment completed successfully"