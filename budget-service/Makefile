PROTO_DIR=proto
PROTO_FILE=budget-service.proto

gen-proto:
	@protoc \
		--go_out=. \
		--go-grpc_out=. \
		$(PROTO_DIR)/$(PROTO_FILE)

DB_URL=postgres://postgres:1@localhost:5432/postgres?sslmode=disable

mig-create:
	@migrate create -ext sql -dir ../migrations -seq create_budgets_table

migrate-up:
	@migrate -path ./migrations -database $(DB_URL) up

migrate-down:
	@migrate -path ./migrations -database $(DB_URL) down

migrate-force:
	@migrate -path=/migrations -database $(DB_URL) verbose -force 1
