# build stage
FROM golang:alpine AS build-env
RUN apk add --no-cache git
WORKDIR $GOPATH/src/gitlab.com/cloudowski/krazy-cow
COPY . .
COPY config /app/config/
COPY web /app/web/

RUN go get -v -d
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /app/goapp

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /app/ .
EXPOSE 8080
USER 1001
ENTRYPOINT [ "/app/goapp" ]
