FROM golang:1.22.2-alpine3.19 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o main .

# Start a new stage from a smaller base image to create a final image
FROM alpine:3.19
COPY --from=builder /app/main /main
CMD ["./main"]
