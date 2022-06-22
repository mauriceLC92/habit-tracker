startdb:
	docker container start habit-postgres

createdb:
	docker run --name habit-postgres -e POSTGRES_PASSWORD=mysecretpassword -e POSTGRES_DB=habit-tracker -p 5432:5432 -d postgres

start-web:
	go run ./cmd/web

migrate-create:
	migrate create -ext sql -dir internal/db/migrations -seq $(filename)

migrate-up:
	migrate -path internal/db/migrations -database "postgresql://postgres:mysecretpassword@localhost:5432/habit-tracker?sslmode=disable" -verbose up

migrate-down:
	migrate -path internal/db/migrations -database "postgresql://postgres:mysecretpassword@localhost:5432/habit-tracker?sslmode=disable" -verbose down

generate:
	sqlc generate