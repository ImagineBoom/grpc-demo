# grpc-demo

## 生成基础服务代码

`cd grpc-demo/four_methods/src`

`protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative src/protobuf/four_methods.proto`

`protoc -I src/protobuf/ src/protobuf/four_methods.proto --dart_out=grpc:lib/protobuf`

## 下载对应依赖包
### Dart plugin for the protocol compiler:

Install the protocol compiler plugin for Dart (protoc-gen-dart) using the following command:

`dart pub global activate protoc_plugin`

Update your PATH so that the protoc compiler can find the plugin:

`export PATH="$PATH:$HOME/.pub-cache/bin"`

### Go plugins for the protocol compiler:

Install the protocol compiler plugins for Go using the following commands:

`go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28`
`go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2`

Update your PATH so that the protoc compiler can find the plugins:

`export PATH="$PATH:$(go env GOPATH)/bin"`