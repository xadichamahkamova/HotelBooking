PROTO_DIR=protos

gen-proto:
	@protoc \
		--go_out=. \
		--go-grpc_out=. \
		$(PROTO_DIR)/booking-service/booking.proto \
		$(PROTO_DIR)/user-service/user.proto \
		$(PROTO_DIR)/hotel-service/hotel.proto \
		$(PROTO_DIR)/notification-service/notification.proto 
		
SWAGGER := ~/go/bin/swag
SWAGGER_DOCS := docs
SWAGGER_INIT := $(SWAGGER) init -g ./internal/http/api.go -o $(SWAGGER_DOCS)

swag-gen:
	$(SWAGGER_INIT)
