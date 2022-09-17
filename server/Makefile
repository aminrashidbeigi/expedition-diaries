generatesql: ## compile sqls
	cd backend/storage && sqlc generate

test: ## to run tests
	go test -v -timeout 30s ./...

build:
	go build -o="history-travels"