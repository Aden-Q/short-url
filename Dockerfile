# syntax=docker/dockerfile:1
# minimize the image size with multi-stage builds
FROM docker.io/golang:1.21.4 AS build-stage

WORKDIR /

COPY . ./

RUN go mod download \
    && CGO_ENABLED=0 GOOS=linux go build -o /short-url

# deploy the application binary into a lean image
FROM docker.io/alpine:edge AS build-release-stage

WORKDIR /

COPY --from=build-stage /short-url /short-url
# the env file holding settings for the application
COPY --from=build-stage /.env /.env


RUN apk update \
    && apk add --no-cache curl sudo

EXPOSE 8080/tcp

USER root:root

ENTRYPOINT ["/short-url"]
