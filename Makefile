DB_URL=postgresql://root:secret@localhost:5435/gqlgen?sslmode=disable

gl_init:
	go run github.com/99designs/gqlgen init && go mod tidy

gl_generate:
	go run github.com/99designs/gqlgen generate --verbose && go mod tidy

new_migration:
	migrate create -ext sql -dir postgres/migration -seq $(name)

migrateup:
	sudo migrate -path postgres/migration -database "$(DB_URL)" -verbose up

.PHONY: gl_init gl_generate new_migration migrateup