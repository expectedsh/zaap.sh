FILES=scheduler.proto
RUBY_OUT_DIR=../zaap-apiserver/lib/protocol

protoc \
  -I. \
  -I/usr/local/include \
  -I$GOPATH/src \
  --go_out=plugins=grpc:../zaap-scheduler/pkg/protocol \
  $FILES

grpc_tools_ruby_protoc \
  -I. \
  --ruby_out=$RUBY_OUT_DIR \
  --grpc_out=$RUBY_OUT_DIR \
  $FILES
