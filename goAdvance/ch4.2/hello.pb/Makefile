generate:
#	protoc --go_out=. hello.proto
#	protoc --go_out=plugins=grpc:. hello.proto
	protoc --go-netrpc_out=plugins=netrpc:. hello.proto

clean:
	-rm *.pb.go
