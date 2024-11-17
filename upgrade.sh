#!/bin/bash

# Pull the latest changes from the repository
git pull

# get the
docker compose pull

# Restart the containers
docker compose up -d

# Remove the old images
docker image prune -f
