##### STAGE 0 #####
FROM golang:1.14.3-alpine3.11 as base

RUN apk update && apk add --no-cache git gcc libc-dev
RUN go get -u github.com/golang/protobuf/proto && go get -u github.com/golang/protobuf/protoc-gen-go
RUN go get -u github.com/micro/protoc-gen-micro
RUN GO111MODULE=on go get github.com/micro/micro/v2

WORKDIR /app
COPY ./go.mod /app/
COPY ./go.sum /app/
RUN go mod download
##### END STAGE 0 #####

##### STAGE 1 #####
FROM base as dev
COPY . /app/
RUN cd /app/service/notification && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./service-notification .
##### END STAGE 1 #####

##### STAGE 2 #####
FROM alpine:latest as prod
RUN apk --no-cache add ca-certificates
RUN mkdir config
COPY --from=dev /app/service/notification/config  /config
COPY --from=dev /app/service/notification/service-notification  /
CMD ./service-notification
##### END STAGE 2 #####
