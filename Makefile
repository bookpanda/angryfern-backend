dev:
	go run cmd/main.go

prod:
	docker-compose -f docker-compose.prod.yaml up

lint:
	golangci-lint run

inspect: 
	atlas schema inspect \
	--url "postgres://root:1234@localhost:5432/angryfern_db?search_path=public&sslmode=disable" \
	--format "{{ sql . }}" > ./migrations/init.sql

FILE := init.sql
migrate:
	atlas schema apply \
	--url "postgres://root:1234@127.0.0.1:5432/angryfern_db?&sslmode=disable" \
	--to "file://./migrations/$(FILE)" \
	--dev-url "docker://postgres/15"