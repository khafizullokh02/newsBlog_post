PSQL_URL=postgres://root:secret@localhost:5432/newsBlog_post?sslmode=disable

postgres:
	docker run --name postgres14 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it postgres14 createdb --username=root --owner=root newsBlog_post

dropdb:
	docker exec -it postgres14 dropdb newsBlog_post

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/newsBlog_post?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/newsBlog_post?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/newsBlog_post?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/newsBlog_post?sslmode=disable" -verbose down 1

execdb:
	docker exec -it postgres14 psql -U root -d newsBlog_post

sqlc:
	sqlc generate

cleandb:
	docker exec -it postgres14 psql -c "DROP SCHEMA public CASCADE; CREATE SCHEMA public;" ${PSQL_URL}

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/khafizullokh02/newsBlog_post/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 execdb sqlc cleandb test server mock