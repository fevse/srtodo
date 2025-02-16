include .env
BIN := "./bin/srtodo"

build:
	go build -o $(BIN) ./cmd/srtodo

run: build
	$(BIN)

migrate:
	goose -dir db/migrations postgres "${DBTYPE}://${DBUSER}:${DBPASS}@${DBADDR}/${DBNAME}?sslmode=disable" up

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.57.2

lint: install-lint-deps
	golangci-lint run ./...