ifneq (,$(wildcard ./.env))
    include .env
    export
endif

init-macos:
	brew install go
	brew install sqlc
	brew install golang-migrate
	make migrate-up

migrate-up:
	migrate -database sqlite://$(DB_URL) -path "./db/migrations" up

migrate-down:
	migrate -database sqlite://$(DB_URL) -path "./db/migrations" down

generate-db:
	sqlc generate
