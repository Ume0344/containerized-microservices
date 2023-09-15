FROM golang:1.20.5-alpine3.18

# Set working directory
WORKDIR /microservice1/

# Copy the microservice1
COPY . /microservice1/

# Build the go application and output to file docker-microservice1
RUN CGO_ENABLED=0 GOOS=linux go build -o docker-microservice1

# Run the microservice1.
CMD ["./docker-microservice1"]
