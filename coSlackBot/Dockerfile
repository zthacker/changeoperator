FROM golang:alpine as builder

# Set current working directory
WORKDIR /app

# Copy the source code
COPY . .

# Note here: CGO_ENABLED is disabled for cross system compilation
# Build the application.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/main .

## Here we'll create out distro-image; if you use scratch, be sure to copy over the certificates; otherwise Slack will not work
FROM alpine:3.12
COPY --from=builder /app/bin/main .

## Expose 9500 on the container for Slack to hit it
## Note that you'll need to setup something like NGROK if you're doing this locally
EXPOSE 9500

# Run executable
CMD ["./main"]