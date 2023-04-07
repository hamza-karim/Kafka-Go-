# Use an official Golang runtime as a parent image
FROM golang:1.16

# Set the working directory to /enricher
WORKDIR /enricher

# Copy the current directory contents into the container at /enricher
COPY . /enricher

# Install the GCC compiler
RUN apt-get update && apt-get install -y gcc

# Install the kafka library
RUN go get -u github.com/confluentinc/confluent-kafka-go/kafka

# Install the protobuf library
RUN go get -u github.com/golang/protobuf/proto

# Install the JSONpb library
RUN go get -u github.com/golang/protobuf/jsonpb

# Build the enricher binary
RUN go build -o main .

# Run enricher when the container starts
CMD ["./main"]
