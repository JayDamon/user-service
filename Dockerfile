# FROM golang:1.18-alpine as builder

# RUN mkdir /app

# COPY . /app

# WORKDIR /app

# RUN CGO_ENABLED=0 go build -o userService ./cmd/main

# RUN chmod +x /app/userService

FROM alpine:latest

RUN mkdir /app

COPY userService /app
# COPY --from=builder /app/userService /app

CMD ["/app/userService"]