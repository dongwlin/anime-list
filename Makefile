.PHONY: build
build:
	go run ./scripts/build

.PHONY: wire
wire:
	wire gen cmd/wire.go

.PHONY: sqlc
sqlc:
	sqlc generate -f sqlc.yml

.PHONY: mock
mock:
	mockgen -package mock -destination internal/store/mock/store.go github.com/dongwlin/anime-list/internal/store Store
