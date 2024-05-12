# Database migrations

This project uses go migrate for it's postgreSql databse migrations
check [go migrate](https://github.com/golang-migrate/migrate) to view go migrate documentation

## Creating migration files

To create a migration file run:

```
  migrate create -ext psql -dir db/migrations -seq MIGRATION_NAME
```

## Testing migrations
To test migration run the up and down of the migrations

run:
```
 migrate -path db/migrations  -database "postgresql://postgres:YOURDATABASEPASSWORD@localhost:5432/DATABASE_NAME?sslmode=disable" -verbose  up 
```
```
 migrate -path db/migrations  -database "postgresql://postgres:YOURDATABASEPASSWORD@localhost:5432/DATABASE_NAME?sslmode=disable" -verbose  down 
```