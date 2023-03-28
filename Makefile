run:
	go run cmd/server/main.go

test:
	go test -v ./... -race

migrate:
	go run cmd/migrate/main.go

rollback:
	go run cmd/rollback/main.go