PROTO_DIR=proto
PROTO_FILE=notification-service.proto

gen-proto:
	@protoc \
		--go_out=. \
		--go-grpc_out=. \
		$(PROTO_DIR)/$(PROTO_FILE)