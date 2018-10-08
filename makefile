.PHONY: build local-db

build:
	go build -o main ./cmd/main.go

local-db:
	@docker-compose down
	@docker-compose up -d

run: build
	./main