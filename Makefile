postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root notes

dropdb:
	docker exec -it postgres12 dropdb notes

migrateup:
	migrate -path database/migration -database "postgresql://root:secret@localhost:5432/notes?sslmode=disable" -verbose up

migratedown:
	migrate -path database/migration -database "postgresql://root:secret@localhost:5432/notes?sslmode=disable" -verbose down

migrateversion:
	migrate -path database/migration -database "postgresql://root:secret@localhost:5432/notes?sslmode=disable" version

.PHONY: postgres createdb dropdb migrateup migratedown migrateversion