build:
	protoc -I. --go_out=plugins=grpc:$(shell pwd)  proto/consignment/consignment.proto

micro:
	protoc -I. --go_out=plugins=micro:$(shell pwd)  proto/consignment/consignment.proto