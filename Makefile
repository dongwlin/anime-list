.PHONY: build
build:
	go run ./scripts/build

.PHONY: sqlc
sqlc:
	sqlc generate -f ./third_party/sqlc/sqlc.yml

.PHONY: mock
mock:
	mockgen -package mock -destination internal/mock/store.go github.com/dongwlin/anime-list/internal/db Store
