PROTO_DIR=proto
PROTO_FILE=hotel-service.proto

MIGRATE_FLAGS=-path migrations -database "postgres://postgres:1@localhost:5432/postgres?sslmode=disable"

gen-proto:
	@protoc \
		--go_out=. \
		--go-grpc_out=. \
		$(PROTO_DIR)/$(PROTO_FILE)

mig-create-hotels:
	@migrate create -ext sql -dir migrations -seq create_hotels_table

mig-create-rooms:
	@migrate create -ext sql -dir migrations -seq create_rooms_table

mig-up:
	@migrate $(MIGRATE_FLAGS) up

mig-down:
	@migrate $(MIGRATE_FLAGS) down
	
mig-force:
	@migrate $(MIGRATE_FLAGS) force 1
