list1:
	grpcurl -plaintext localhost:1234 list
	
list2:
	grpcurl -plaintext localhost:1234 list helloservice.HelloService
	
describe1:
	grpcurl -plaintext localhost:1234 describe helloservice.HelloService
	
describe2:
	grpcurl -plaintext localhost:1234 describe helloservice.String
	
run1:
	grpcurl -plaintext -d '{"value": "gopher"}' localhost:1234 helloservice.HelloService/Hello
	
run2:
	grpcurl -plaintext -d @ localhost:1234 helloservice.HelloService/Channel