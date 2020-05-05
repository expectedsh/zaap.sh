PROTO_FILES=\
	proto/runner.proto
PROTO_FLAGS=-Iproto -I/usr/local/include

make-protocol:
	protoc $(PROTO_FLAGS) --go_out=plugins=grpc:./pkg/protocol $(PROTO_FILES)
	protoc $(PROTO_FLAGS) --go_out=plugins=grpc:../zaap-services/pkg/protocol $(PROTO_FILES)
