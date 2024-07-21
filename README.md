## Description

This is go-esb-test.

## Tech Stack

- Golang : https://github.com/golang/go
- MySQL (Database) : https://github.com/mysql/mysql-server
- Docker: https://www.docker.com/

## Framework & Library

- GoFiber (HTTP Framework) : https://github.com/gofiber/fiber
- GORM (ORM) : https://github.com/go-gorm/gorm
- Viper (Configuration) : https://github.com/spf13/viper
- Golang Migrate (Database Migration) : https://github.com/golang-migrate/migrate
- Go Playground Validator (Validation) : https://github.com/go-playground/validator
- Logrus (Logger) : https://github.com/sirupsen/logrus

## Configuration

All configuration is in `config.json` file.
Please check and update it with your database configuration.

## Command
All command is in `Makefie` file.
Please check and update it with your configuration / command.

## Collection Postman API
- Postman : https://elements.getpostman.com/redirect?entityId=2255661-a91d7465-25aa-4667-8f2e-296adaceea13&entityType=collection

## Install GO Modules

```shell
go mod tidy
```

## Database Migration

All database migration is in `db/migrations` folder.

### Create Migration

```shell
make migrate_create file=ceate_table
```

### Create DB
```shell
make createdb
```

### Create Dummy Data (SEED)
```shell
make seed
```

### Run Migration

```shell
make migrate_up
```

## Run Application

### Run unit test

```bash
make test
```

### Run web server

This addition specifies that users should update the `database config` in `config.json` if they are not using Docker, ensuring that the application can connect to their local MySQL instance.

```bash
make run
```

Using Docker
To run the application using Docker:

```bash
make run_mysql
make server
```