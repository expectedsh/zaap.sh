protoc \
  -I. \
  -I/usr/local/include \
  -I$GOPATH/src \
  --go_out=plugins=grpc:. \
  --include_imports \
  --include_source_info \
  app.proto
