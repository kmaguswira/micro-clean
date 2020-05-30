##### STAGE 1 #####
FROM golang:1.14.3-alpine3.11 as dev

RUN apk update && apk add --no-cache git gcc libc-dev
RUN go get -u github.com/golang/protobuf/proto && go get -u github.com/golang/protobuf/protoc-gen-go
RUN go get -u github.com/micro/protoc-gen-micro
RUN GO111MODULE=on go get github.com/micro/micro/v2

WORKDIR /app
COPY . /app/
RUN go mod download

##### BUILD APPS OR SERVICES #####
RUN cd /app/service/account && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./service-account
##### END STAGE 1 #####

##### STAGE 2 #####
FROM alpine:latest as prod
RUN apk --no-cache add ca-certificates
RUN mkdir config
COPY --from=dev /app/service/account/config  /config
COPY --from=dev /app/service/account/service-account  /
CMD ./service-account
##### END STAGE 2 #####
