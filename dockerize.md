# Dockerization Guide
This guide will walk you through the process of dockerizing this application.

## Quick Start
Build and run the Docker image.

```bash
# Build the Docker image with the Dockerfile
docker build -t heyapi .

# Run the Docker image
docker run -p 8080:8080 heyapi
```

**Access the application at http://localhost:8080.**

## Verify the running container.
Here are some convenient commands for troubleshooting:

```bash
# List all running containers
docker ps -a

# List all loaded images
docker images

# To remove the container
docker rm <container_id>

# To remove the image
docker image rm heyapi
```

## To save the docker image
Save the image as a tar file after building it.

```bash
docker save -o heyapi.tar heyapi
```

## To load the docker image
Load the image from a tar file on another machine.

```bash
docker load -i heyapi.tar
```
