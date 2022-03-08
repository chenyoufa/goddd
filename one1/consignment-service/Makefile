build:
	protoc -I. --go_out=plugins=grpc:$(shell pwd)  proto/consignment/consignment.proto 
	GOOS=linux GOARCH=amd64 go build    
	docker build -t consignment-service .
run:
	docker run -p 50051:50051 consignment-service

micro:
	protoc -I. --go_out=plugins=micro:$(shell pwd)  proto/consignment/consignment.proto


	/home/fage/workspace/go/thrgoweb/Dockerfile