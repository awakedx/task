# Task 1 CRUD HTTP API with Basic Authentication
 - **Username:** `admin`
 - **Password:** `password`

## Used technologies

- **Golang**: Standard library net/http
- **PostgreSQL**: Free and open-source relational database
- **Docker**: Docker allows you to run an app in containers, that provides isolation, simplified deployment and control of dependencies

## Dependencies

- **github.com/go-playground/validator/v10**: For validating struct fields
- **github.com/jackc/pgx**: PostgreSQL driver and toolkit for Go
- **github.com/spf13/viper**: Config managing
- **github.com/golang-migrate/migrate**: For migration database

## Build And Run
### Start app
```bash
make up
```
### Build and start in docker

```bash
make docker-build
```

# Task 2 Bench of concatenation function
### conc_test.go file located in cmd/task_2
```bash
make bench
```
