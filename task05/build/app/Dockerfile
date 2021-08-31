FROM alpine:latest AS u-shylianok-task05-app

RUN apk add --no-cache \
        libc6-compat

RUN mkdir -p /app
WORKDIR /app

ADD task05 /app
EXPOSE 8000
CMD ["./task05"]