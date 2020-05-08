# Microservice Skeleton with Micro and Gin

This repository is a boilerplate to start build a microservice with **Micro** and **Gin**. The basic features of this boilerplate are **ACL** (Access Control List), **Notification** (Email Notification with **SendGrid**) and **File Manager** (will be integrate with a CDN). Each feature created in a service, and each service is not depend with the other (loose coupling). The application can access a service via **gRPC** and also can publish an event to do something in a service.


## Getting Started
If you want to start with this repository you must install the following requirement:

	-	docker run --name mysql -p 3306:3306 -d mysql

	-	HOMEBREW_NO_AUTO_UPDATE=1 brew install protobuf

	-	HOMEBREW_NO_AUTO_UPDATE=1 brew install consul

	-	go get -u github.com/golang/protobuf/{proto,protoc-gen-go}

	-	go get github.com/micro/micro/v2

	-	go get -u github.com/micro/protoc-gen-micro
	
	-	micro cli

## Starting The Application
-	Make sure **consul** is running with `consul agent -dev`
-	Build the proto file of each service with `make proto` in the root of directory of each service
-	Run the **DB seeder** of each service with `go run main.go` in the `repositories/seeder` of each service
-	Run the service of with `go run main.go` in the root of directory of each service
-	Run the application in the `app` folder with `go run main.go`
## Architecture
Each service is created with `micro`, and with some simple folder structuring, it can be clean architecture. Each service follow this folder template:
```
+-service
|
+--- ....
|
+---account
	+ application
		+ usecase (business logic)
		+ repositories (interface of repository)
		+ external_service (interface of external service)
		+ global (constant and other shared input)
	+ config
	+ domain 
	+ external_service (third party service)
	+ handler 
	+ proto (contract)
	+ repositories
		+ entity (db model)
		+ seeder (db seeder)
	+ subscriber (subcriber handler)
	+ utils (helper)	
```
![Architecture](https://raw.githubusercontent.com/kmaguswira/micro-clean/master/doc/image/arch.png)

## Request Flow

![Request Flow](https://raw.githubusercontent.com/kmaguswira/micro-clean/master/doc/image/req-flow.png)

## Postman Collections
[Account Service](https://github.com/kmaguswira/micro-clean/blob/master/app/web/Account.postman_collection.json)
[Notification Service](https://github.com/kmaguswira/micro-clean/blob/master/app/web/Notification.postman_collection.json)
[File Service](https://github.com/kmaguswira/micro-clean/blob/master/app/web/File.postman_collection.json)

## TODO
- [ ] Unit Testing
- [ ] Dockerize
- [ ] Error Handling
- [ ] Refactor some codes
