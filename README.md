# Golang Boilerplate

## Dependencies

Golang 1.17 & Postgres

- Gin
- Godotenv
- Golang-Migrate
- Sqlx
- Sqlmock

## How to Test

```
make test
```

## How to Run

1. Initialize environment variable into `.env` file

2. Prepare database

3. Migrate table schema

   ```
   make migrate
   ```

4. Run

   ```
   make run
   ```

## Add new migration

1. Ensure `golang-migrate` CLI has been installed on your machine
2. Run

   ```
   migrate create -ext sql -dir migrations -seq <changes_filename>
   ```

3. Migration file will be created in `/migrations` dir
