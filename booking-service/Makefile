PROTO_DIR=proto
PROTO_FILE=booking-service.proto

MIGRATE_FLAGS=-path migrations -database "postgres://postgres:1@localhost:5432/postgres?sslmode=disable"

gen-proto:
	@protoc \
		--go_out=. \
		--go-grpc_out=. \
		$(PROTO_DIR)/$(PROTO_FILE)

mig-create:
	@migrate create -ext sql -dir migrations -seq create_bookings_table

mig-up:
	@migrate $(MIGRATE_FLAGS) up

mig-down:
	@migrate $(MIGRATE_FLAGS) down
	
mig-force:
	@migrate $(MIGRATE_FLAGS) force 1
