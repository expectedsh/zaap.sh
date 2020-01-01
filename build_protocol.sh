protoc \
  -Iprotocol \
  -I/usr/local/include \
  -I$GOPATH/src \
  -I/Users/remi/Documents/Go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:protocol \
  --descriptor_set_out=protocol/app.pb \
  --include_imports \
  --include_source_info \
  protocol/app.proto
