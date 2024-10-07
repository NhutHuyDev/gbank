postgres:
	docker run --name sgbankpg12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb: 
	docker exec -it sgbankpg12 createdb --username=root --owner=root sgbank

dropdb:
	docker exec -it sgbankpg12 dropdb sgbank

migrateup:
	migrate -path db/migration -database "postgres://root:secret@127.0.0.1:5432/sgbank?sslmode=disable" -verbose up 

migratedown:
	migrate -path db/migration -database "postgres://root:secret@127.0.0.1:5432/sgbank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test: 
	go test -v -cover ./...

server:
	go run ./cmd/sgbank/main.go

.PHONY: posgres createdb dropdb migrateup migratedown sqlc test server