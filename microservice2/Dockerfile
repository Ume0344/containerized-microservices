FROM golang:1.20.5-alpine3.18

# Set working directory
WORKDIR /microservice2/

# Copy the microservice2
COPY . /microservice2/

# Build the go application and output to file docker-microservice2
RUN CGO_ENABLED=0 GOOS=linux go build -o docker-microservice2

# Run the microservice2
CMD ["./docker-microservice2"]
