aa:
	cd ../../.. && pwd

gen:
	protoc -I=. -I=/Users/xujinzhao/go/pkg/mod --go_out=plugins=grpc:. helloworld.proto
	protoc -I=. -I=/Users/xujinzhao/go/pkg/mod --govalidators_out=. helloworld.proto

clean:
	-rm *.pb.go
