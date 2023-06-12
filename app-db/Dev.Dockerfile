FROM golang:1.19-alpine
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64;

WORKDIR /app/

# Copy go mod files
COPY go.mod go.sum \
     /app/

RUN go mod download

COPY .  /app/


RUN go build -o /tmp/garination-database


FROM alpine:latest
RUN mkdir -p /usr/bin/
COPY --from=0 /tmp/garination-database /usr/bin/garination-database
ENTRYPOINT ["/usr/bin/garination-database"]