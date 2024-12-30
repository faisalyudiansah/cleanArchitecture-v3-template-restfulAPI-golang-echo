# Build stage
FROM golang:1.22.9 AS build-stage

WORKDIR /app

# Copy go module files first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Create logs directory and set permissions
RUN mkdir -p /app/public/logs && \
    chmod 755 /app/public/logs

# Build the applications
RUN CGO_ENABLED=0 GOOS=linux go build -o /web cmd/web/main.go && \
    CGO_ENABLED=0 GOOS=linux go build -o /worker cmd/worker/main.go

# Release stage
FROM gcr.io/distroless/static-debian12 AS build-release-stage

WORKDIR /

# Copy .env file for configuration
COPY .enva /.env

# Set default environment variables
ENV PORT=8011 \
    TIMEZONE=Asia/Jakarta \
    TIMEOUT_GRACEFUL_SHUTDOWN=30 \
    MYSQL_READ_HOST= \
    MYSQL_READ_USER= \
    MYSQL_READ_PASS= \
    MYSQL_READ_DBNAME=server_v3_db \
    MYSQL_READ_PORT=4000 \
    MYSQL_READ_USE_TLS=true \
    MYSQL_WRITE_HOST= \
    MYSQL_WRITE_USER= \
    MYSQL_WRITE_PASS= \
    MYSQL_WRITE_DBNAME=server_v3_db \
    MYSQL_WRITE_PORT=4000 \
    MYSQL_WRITE_USE_TLS=true \
    DB_MAX_OPEN_CONNS=10 \
    DB_MAX_IDLE_CONNS=10 \
    DB_CONN_MAX_LIFETIME=300 \
    DB_LOG_LEVEL=info \
    RABBITMQ_HOST= \
    RABBITMQ_PORT=5672 \
    RABBITMQ_USER= \
    RABBITMQ_PASSWORD= \
    RABBITMQ_VHOST=/ \
    RABBITMQ_HEARTBEAT_INTERVAL=10 \
    RABBITMQ_RETRY_CONNECT_INTERVAL=10 \
    RABBITMQ_MAX_RETRY_CONNECT=4

# Copy binaries and public directory from build stage
COPY --from=build-stage /web /web
COPY --from=build-stage /worker /worker
COPY --from=build-stage /app/public /public

EXPOSE ${PORT}

ENTRYPOINT [ "/web" ]