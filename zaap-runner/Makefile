PROTO_FILES=pkg/protocol/runner.proto
PROTO_FLAGS=-Ipkg/protocol -I/usr/local/include

make-protocol:
	protoc $(PROTO_FLAGS) --go_out=plugins=grpc:./pkg/protocol $(PROTO_FILES)
