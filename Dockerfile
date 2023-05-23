# Use a Go image as the base
FROM golang:latest
RUN apt-get update && apt-get install -y nano

# Set the working directory
WORKDIR /app

# Copy the Go source code into the container
COPY . .

# Build the Go application
RUN go build -o app

# Set the command to run when the container starts
CMD [ "./app" ]
