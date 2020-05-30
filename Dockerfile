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
RUN cd /app/app/web && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./app-web
RUN cd /app/service/account && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./service-account
RUN cd /app/service/file && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./service-file
RUN cd /app/service/notification && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./service-notification

##### END STAGE 1 #####

##### STAGE 2 #####
FROM alpine:latest as prod
RUN apk --no-cache add ca-certificates
RUN mkdir config
COPY --from=dev /app/app/web/config  /config
COPY --from=dev /app/app/web/app-web  /config
CMD ./app-web
##### END STAGE 2 #####
