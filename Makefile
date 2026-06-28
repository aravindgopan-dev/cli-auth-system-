DB_URL=postgres://postgres:postgres@localhost:5432/auth_db?sslmode=disable
MIGRATION_PATH=./migration

migrate-up:
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" up

run:
	sudo docker compose run --build app

clean:
	sudo docker compose down -v