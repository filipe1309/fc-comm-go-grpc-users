```sh
go mod init github.com/codeedu/fc2-grpc

sudo apt install protobuf-compiler
protoc --version

go get google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go

protoc --proto_path=proto proto/*.proto --go_out=pb

go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb
```
