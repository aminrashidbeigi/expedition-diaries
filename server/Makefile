generatesql: ## compile sqls
	cd storage && sqlc generate

test: ## to run tests
	go test -v -timeout 30s ./...

build:
	go build -o="expedition-diaries" ./cmd
