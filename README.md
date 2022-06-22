# Habit tracker

The aim of this project is to produce a Go package and accompanying command-line tool that will help users track and establish a new habit, by reporting their current streak.

## MVP tasks

- As a user I want to create a habit to track
- As a user I want to be able to log consecutive days for my habit
- As a user I want to see my current streak (if any)

# Database

The database is postgresql, both migrations and queries are vanilla SQL

### Start postgresql for testing

    `make postgres`
    `make createdb`

### Migrate

[golang-migrate](https://github.com/golang-migrate/migrate) reads migrations from `store/migration` older and applies them in the correct order.

- Create a new migration by running `migrate create -ext sql -dir store/migration -seq [NAME]`
  Replace NAME with something like init_schema, or the jira ticket number.

- Edit the generated `store/migration/___up.sql and down.sql` files

- run the migration

  `make migrateup`

### SQLC

[sqlc](https://sqlc.dev/) is used to compile SQL to type-safe Go

Initialize sqlc, creating the config file `sqlc.yaml`
`sqlc init`

