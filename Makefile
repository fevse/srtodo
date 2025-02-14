include .env
BIN := "./bin/srtodo"

build:
	go build -o $(BIN) ./cmd/srtodo

run: build
	$(BIN)

migrate:
	goose -dir db/migrations postgres "${DBTYPE}://${DBUSER}:${DBPASS}@${DBADDR}/${DBNAME}?sslmode=disable" up