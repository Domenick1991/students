.PHONY: generate-api
generate-api:
	@./scripts/generate.sh

.PHONY: up
up:
	podman-compose up -d

.PHONY: down
down:
	podman-compose down

