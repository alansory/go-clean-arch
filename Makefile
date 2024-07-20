# Pull the MySQL 5.7 Docker image
pull_mysql:
	docker pull mysql:5.7

# Stop and remove the existing MySQL container
stop_remove_mysql:
	docker stop mysql-local || true
	docker rm mysql-local || true

# Create a new Docker volume
create_volume:
	docker volume create mysql_data_local || true

# Remove the Docker volume
remove_volume:
	docker volume rm mysql_data_local || true

# Run the MySQL 5.7 Docker container
run_mysql: stop_remove_mysql remove_volume create_volume
	docker run --name mysql-local -e MYSQL_ROOT_PASSWORD=secret -p 3307:3306 -v mysql_data_local:/var/lib/mysql -d mysql:5.7

# Create a new database
createdb:
	docker exec -it mysql-local mysql -u root -psecret -e "CREATE DATABASE go_esb"

# Drop the existing database
dropdb:
	docker exec -it mysql-local mysql -u root -psecret -e "DROP DATABASE go_esb"

# Create a new migration
migrate_create:
	migrate create -ext sql -dir db/migrations $(file)

# Apply all migrations
migrate_up:
	migrate -database "mysql://root:secret@tcp(127.0.0.1:3307)/go_esb" -path db/migrations up

# Rollback the last migration
migrate_down:
	migrate -database "mysql://root:secret@tcp(127.0.0.1:3307)/go_esb" -path db/migrations down

# Run server
run_server:
	go run cmd/web/main.go 

.PHONY: pull_mysql stop_remove_mysql create_volume remove_volume run_mysql createdb dropdb migrate_create migrate_up migrate_down run_server
