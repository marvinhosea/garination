FROM golang:1.20-alpine

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64;

WORKDIR /app/

# Copy go mod files
COPY go.mod go.sum \
     /app/

RUN go mod download

COPY .  /app/


RUN mkdir -p /usr/bin/
RUN go build -o /usr/bin/garination-gateway

FROM alpine:latest
RUN mkdir -p /usr/bin/
COPY --from=0 /usr/bin/garination-gateway /usr/bin/garination-gateway

ENTRYPOINT ["/usr/bin/garination-gateway"]