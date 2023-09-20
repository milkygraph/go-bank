postgres:
	docker run --name postgres --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=kurero17 -d postgres:12-alpine

docker_server:
	docker run --name go-bank --network bank-network -p 8080:8080 -e DB_SOURCE="postgresql://root:kurero17@postgres:5432/simple_bank?sslmode=disable" go-bank:latest

remove-postgres:
	docker rm -f postgres

createdb:
	docker exec -it postgres createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:kurero17@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:kurero17@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:kurero17@localhost:5432/simple_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:kurero17@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen --package mockdb --destination db/mock/store.go github.com/milkygraph/simplebank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown server mock
