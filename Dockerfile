# Build stage
FROM golang:1.22.5-alpine AS builder
WORKDIR /app
COPY . .
RUN go build
RUN apk --no-cache add curl

# Run stage
FROM golang:1.22.5-alpine
WORKDIR /app
COPY --from=builder /app/statusneo .


EXPOSE 8001
CMD [ "/app/statusneo" ]

