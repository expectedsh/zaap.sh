PROTO_FILES=\
	proto/application_events.proto \
	proto/runner_events.proto
PROTO_FLAGS=-Iproto -I/usr/local/include

make-protocol:
	protoc $(PROTO_FLAGS) --go_out=./pkg/protocol $(PROTO_FILES)
