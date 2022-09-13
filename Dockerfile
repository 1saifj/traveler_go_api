# Start from golang base image
FROM golang:1.19.1-alpine3.16

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base

# Setup folders
WORKDIR /go/src

# Copy the source from the current directory to the working Directory inside the container
COPY . ./bff

# Change Workdir  for bff
WORKDIR /go/src/bff


ENV GO111MODULE=on
# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

EXPOSE 8080

# Build the Go app
RUN go build -o /build

# Run the executable
CMD [ "/build" ]
