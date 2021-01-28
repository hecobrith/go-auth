
```shell
$ sudo apt install -y protobuf-compiler
$ protoc --version  # Ensure compiler version is 3+


sudo apt install golang-goprotobuf-dev

# go install google.golang.org/protobuf/cmd/protoc-gen-go  alternative package

protoc --go_out=. *.proto
# run it on the 


sudo protoc -I protofiles/ protofiles/user.proto --go_out=plugins=grpc:protofiles
```
