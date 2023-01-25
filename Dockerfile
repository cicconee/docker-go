# syntax=docker/dockerfile:1

FROM golang:1.19-alpine

WORKDIR /app 

# COPY go.mod ./
# COPY go.sum ./
# RUN go mod download

# COPY *.go ./

COPY . .
RUN go mod download 

RUN go build -o /docker-go 

CMD ["/docker-go"]