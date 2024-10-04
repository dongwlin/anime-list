.PHONY: build
build:
	go run ./scripts/build

.PHONY: ent
ent:
	go generate ./internal/ent
