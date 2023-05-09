migrate-up:
	migrate -database sqlite3://bank.db -path internal/db/migrations up

migrate-down:
	migrate -database sqlite3://bank.db -path internal/db/migrations down

migrate-create: $(MIGRATE)
	@ read -p "migration name: " Name; \
	migrate create -ext sql -dir internal/db/migrations -seq $${Name}

test:
	go test ./... -v

.PHONY: migrate-up migrate-down migrate-create test