postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=mysecretpassword -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root recipes_db

dropdb:
	docker exec -it postgres12 dropdb recipes_db

migrateup:
	migrate -path datasources/postgres/recipes_db/db_migrations -database "postgresql://root:mysecretpassword@localhost:5432/recipes_db?sslmode=disable" -verbose up

migratedown:
	migrate -path datasources/postgres/recipes_db/db_migrations -database "postgresql://root:mysecretpassword@localhost:5432/recipes_db?sslmode=disable" -verbose down
	
sqlc:
	sqlc generate
	
	.PHONY postgres createdb dropdb migrateup migratedown sqlc