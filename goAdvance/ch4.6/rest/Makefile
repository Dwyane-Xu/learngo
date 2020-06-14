run:
	@go build -o a.out && ./a.out
	-@rm ./a.out

gen:
	protoc \
		-I=. -I=/Users/xujinzhao/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.14.6/third_party/googleapis \
		--go_out=plugins=grpc:. helloworld.proto
	protoc \
		-I=. -I=/Users/xujinzhao/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.14.6/third_party/googleapis \
		--grpc-gateway_out=. helloworld.proto
	protoc \
		-I=. -I=/Users/xujinzhao/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.14.6/third_party/googleapis \
		--swagger_out=. helloworld.proto

clean:
	-rm *.pb.go
