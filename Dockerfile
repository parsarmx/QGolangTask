FROM golang:alpine AS builder

# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Create appuser.
ENV USER=quser
ENV UID=10001

# See https://stackoverflow.com/a/55757473/12429735RUN 
RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"

WORKDIR $GOPATH/Task

COPY . .

# Fetch dependencies.
RUN go mod download
RUN go mod verify
RUN go mod tidy -go=1.19


RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -ldflags '-extldflags "-static"'

RUN which ls
RUN which sh
RUN ls /go/bin

CMD go run cmd/main.go