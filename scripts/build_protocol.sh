protoc \
  -Ipkg/protocol \
  -I/usr/local/include \
  -I$GOPATH/src \
  -I/Users/remi/Documents/Go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:pkg/protocol \
  --descriptor_set_out=pkg/protocol/app.pb \
  --include_imports \
  --include_source_info \
  pkg/protocol/app.proto
