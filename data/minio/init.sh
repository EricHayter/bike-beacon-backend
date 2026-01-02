#!/bin/bash

# Start MinIO server in the background
minio server /data --console-address ":9001" &

# Wait for MinIO to be ready
until curl -f http://localhost:9000/minio/health/live > /dev/null 2>&1; do
  echo "Waiting for MinIO to start..."
  sleep 1
done

echo "MinIO is ready, configuring bucket..."

# Configure mc alias for local MinIO instance using environment variables
mc alias set local http://localhost:9000 ${MINIO_ROOT_USER} ${MINIO_ROOT_PASSWORD}

# Create the images bucket if it doesn't exist
mc mb local/images --ignore-existing

# Set bucket policy to allow public downloads
mc anonymous set download local/images

echo "Bucket 'images' configured for public download access"

# Wait for the MinIO process to keep container running
wait
