FROM golang:1.23.2-alpine

# Set the working directory
WORKDIR /app

# Copy the go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download and install any required Go dependencies
RUN go mod download

# Copy the entire source code to the working directory
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port specified by the PORT enviroment variable 
EXPOSE 3000

# Set the entry point of the container to the executable
CMD ["./main"]