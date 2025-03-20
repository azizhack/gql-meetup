gl_init:
	go run github.com/99designs/gqlgen init && go mod tidy

gl_generate:
	go run github.com/99designs/gqlgen generate && go mod tidy

.PHONY: gl_init gl_generate