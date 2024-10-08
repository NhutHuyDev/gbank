postgres:
	docker run --name sgbankpg12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb: 
	docker exec -it sgbankpg12 createdb --username=root --owner=root sgbank

dropdb:
	docker exec -it sgbankpg12 dropdb sgbank

migrateup:
	migrate -path db/migration -database "postgres://root:secret@127.0.0.1:5432/sgbank?sslmode=disable" -verbose up 

migrateup1:
	migrate -path db/migration -database "postgres://root:secret@127.0.0.1:5432/sgbank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgres://root:secret@127.0.0.1:5432/sgbank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgres://root:secret@127.0.0.1:5432/sgbank?sslmode=disable" -verbose down 1 

sqlc:
	sqlc generate

test: 
	go test -v -cover ./...

server:
	go run ./cmd/sgbank/main.go

mock:
	mockgen -package mockdb -destination internal/infra/db/mock/store.go github.com/NhutHuyDev/sgbank/internal/infra/db Store

.PHONY: posgres createdb dropdb migrateup migrateup1 migratedown migratedown1 sqlc test server