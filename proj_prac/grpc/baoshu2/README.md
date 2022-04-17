 Ref: https://www.cnblogs.com/baoshu/p/13510854.html
 protoc -I . --go_out=plugins=grpc,Mgoogle/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor:. google/api/*.proto