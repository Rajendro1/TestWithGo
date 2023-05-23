# Use a Go image as the base
FROM docker.io/golang:latest
RUN apt-get update && apt-get install -y nano

# Install PostgreSQL
RUN apt-get update && apt-get install -y nano postgresql

# Set the working directory
WORKDIR /app

# Copy the Go source code into the container
COPY . .

# Build the Go application
RUN go build -o app

EXPOSE 8081

# Set the command to run when the container starts
CMD [ "./app" ]
