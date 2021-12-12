FROM golang:1.16-alpine AS base
ENV SERVICE_PORT=80
LABEL go_app.version="v1.0" go_app.author="Bruce Gu"
ADD bin/amd64/http_server_simple /httpserver
EXPOSE 9090
ENTRYPOINT /httpserver