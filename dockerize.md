# Dockerization Guide
Quickly build, run, and manage the Docker image for this application.

## Quick Start
1. **Build the Docker image**:
   ```bash
   docker build -t heyapi .
   ```

2. **Run the container**:
   ```bash
   docker run -p 8080:8080 heyapi
   ```

**Access the app at**: [http://localhost:8080](http://localhost:8080)

---

## Useful Commands
### Verify and Manage Containers
```bash
# List running containers
docker ps

# List all containers (including stopped ones)
docker ps -a

# Remove a container
docker rm <container_id>
```

### Manage Images
```bash
# List all images
docker images

# Remove an image
docker rmi heyapi
```

---

## Save and Load Images
### Save the Image
Export the image to a tar file:
```bash
docker save -o heyapi.tar heyapi
```

### Load the Image
Import the image on another machine:
```bash
docker load -i heyapi.tar
```
