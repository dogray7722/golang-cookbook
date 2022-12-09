postgres:
	docker run --name postgres12 --network cookbook-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=mysecretpassword -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root recipes_db

dropdb:
	docker exec -it postgres12 dropdb recipes_db

migrateup:
	migrate -path datasources/postgres/recipes_db/db_migrations -database "postgresql://root:mysecretpassword@golang-cookbook-1.cbkr6oxrsrka.us-west-2.rds.amazonaws.com:5432/recipes_db" -verbose up

migratedown:
	migrate -path datasources/postgres/recipes_db/db_migrations -database "postgresql://root:mysecretpassword@localhost:5432/recipes_db?sslmode=disable" -verbose down
	
sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server: 
	go run main.go

up:
	docker compose up

down:
	docker compose down
	
mock:
	mockgen -package mockdb -destination datasources/postgres/recipes_db/mock/store.go github.com/golang-cookbook/datasources/postgres/recipes_db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock up down