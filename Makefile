gen:
	@go generate -v ./...

mod:
	@go mod tidy && go mod vendor

lint: ## Run lint
	@./scripts/lint.sh

run:
	@go run ./cmd/services/core/...

build:
	@go build -tags=jsoniter -o ./build/core ./cmd/services/core/...

test:
	@./scripts/test.sh

print:
	@echo $(call args,defaultstring)

migration:
	@migrate create -ext sql -dir databases/mysql/migrations -seq $(name)

migrateup:
	@migrate -source file://databases/mysql/migrations -database "mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${MYSQL_DATABASE}?parseTime=true&charset=utf8mb4" up

migratedown:
	@migrate -source file://databases/mysql/migrations -database "mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${MYSQL_DATABASE}?parseTime=true&charset=utf8mb4" down 1

install-devtools:
	@go get -u github.com/jstemmer/go-junit-report
