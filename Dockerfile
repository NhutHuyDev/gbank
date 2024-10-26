# Build stage
# FROM golang:alpine AS builder
# WORKDIR /app
# COPY . .
# RUN go build -o main ./cmd/sgbank/main.go

# Run stage
# FROM alpine
# WORKDIR /app
# COPY --from=builder /app/main .

# EXPOSE 8080
# CMD ["/app/main"]

FROM golang:alpine 
WORKDIR /app
COPY . .
RUN go build -o main ./cmd/sgbank/main.go

EXPOSE 8080
CMD ["/app/main"]