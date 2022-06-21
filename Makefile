startdb:
	docker run --name habit-postgres -e POSTGRES_PASSWORD=mysecretpassword POSTGRES_DB=habit-tracker -d postgres

generate:
	sqlc generate