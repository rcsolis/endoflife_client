ARG BINARY_NAME="main"
ARG PORT="3000"
# Builder image
FROM golang:1.23.4-bookworm as builder
# Update and install make
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    make && \
    rm -rf /var/lib/apt/lists/*
# Set working directory
WORKDIR /app
# Copy source files
COPY . .
# Build the binary
RUN make build

# Release image
FROM debian:bookworm-slim as release
# Set working directory
WORKDIR /app
# Update and install ca-certificates
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*
# Copy Environment file
COPY --from=builder /app/.env .
# Export environment variables from .env file
RUN set -o allexport; source .env; set +o allexport
# Copy the binary from the builder image
COPY --from=builder /app/bin/ .
# Run the binary
ENV BINARY_NAME=${BINARY_NAME}
ENV PORT=${PORT}
CMD ./app/${BINARY_NAME}