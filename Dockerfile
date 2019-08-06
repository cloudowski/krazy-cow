# build stage
# FROM golang:alpine AS build-env
FROM golang:1.12-buster AS build-env
ARG VERSION
ARG GITCOMMIT
RUN apt-get update && apt-get install --no-install-recommends -y git
WORKDIR $GOPATH/src/gitlab.com/cloudowski/krazy-cow
COPY . .
COPY config /app/config/
COPY web /app/web/

RUN go get -v -d
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s -X main.version=$VERSION -X main.gitCommit=$GITCOMMIT" -o /app/goapp

# final stage
FROM debian:buster-slim
WORKDIR /app
COPY --from=build-env /app/ .
EXPOSE 8080
USER nobody
ENTRYPOINT [ "/app/goapp" ]
