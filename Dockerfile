FROM alpine:latest

RUN mkdir /app


COPY userService /app
COPY /db /db

CMD ["/app/userService"]