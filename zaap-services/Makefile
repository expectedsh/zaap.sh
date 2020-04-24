PROTO_FILES=\
	pkg/protocol/application.proto
PROTO_FLAGS=-Ipkg/protocol -I/usr/local/include

make-protocol:
	protoc $(PROTO_FLAGS) --go_out=./pkg/protocol $(PROTO_FILES)
