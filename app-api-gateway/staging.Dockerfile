FROM golang:1.19-alpine as golang-builder

RUN apk add build-base openssl
RUN apk --update add git ca-certificates
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64;


FROM golang-builder AS app-builder
WORKDIR /app/

# Copy go mod files
COPY go.mod go.sum \
     /app/

RUN go mod download

COPY .  /app/



RUN go build -o /tmp/garination-gateway

FROM app-builder AS prepare-bin

COPY --from=app-builder /tmp/garination-gateway /usr/bin/garination-gateway

ENTRYPOINT ["/usr/bin/garination-gateway"]