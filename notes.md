```sh
go mod init github.com/codeedu/fc2-grpc

sudo apt install protobuf-compiler
protoc --version

go get google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go
-> asdf reshim

protoc --proto_path=proto proto/*.proto --go_out=pb

go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
-> asdf reshim

protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb

go run cmd/server/server.go
```

https://github.com/ktr0731/evans

```sh
go get github.com/ktr0731/evans
go install github.com/ktr0731/evans

https://github.com/ktr0731/evans/issues/150
$ curl -L 'https://github.com/ktr0731/evans/releases/download/0.6.11/evans_linux_amd64.tar.gz' | tar xvzf -
$ file evans
evans: ELF 64-bit LSB executable, x86-64, version 1 (SYSV), dynamically linked, interpreter /lib64/ld-linux-x86-64.so.2, not stripped
$ mv evans /usr/local/bin/evans # This path must be included in $PATH

evans -r repl --host localhost --port 50051
pb.UserService@localhost:50051> service UserService
pb.UserService@localhost:50051> call AddUser

-> go run cmd/server/server.go
-> go run cmd/client/client.go
```
