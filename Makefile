postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres createdb --username=root --owner=root notes

dropdb:
	docker exec -it postgres dropdb notes

migrateup:
	migrate -path database/migration -database "postgresql://root:secret@localhost:5432/notes?sslmode=disable" -verbose up 

migratedown:
	migrate -path database/migration -database "postgresql://root:secret@localhost:5432/notes?sslmode=disable" -verbose down

migrateversion:
	migrate -path database/migration -database "postgresql://root:secret@localhost:5432/notes?sslmode=disable" version

sqlc:
	sqlc generate

server:
	go run main.go

proto:
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
    proto/*.proto

.PHONY: postgres createdb dropdb migrateup migratedown migrateversion server proto