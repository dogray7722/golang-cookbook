postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=mysecretpassword -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root recipes_db

dropdb:
	docker exec -it postgres12 dropdb recipes_db

	.PHONY postgres createdb dropdb