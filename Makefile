.PHONY: up down reset logs

up:
	docker compose up -d

down:
	docker compose down

reset:
	docker compose down -v
	docker compose up -d postgres

logs:
	docker compose logs -f postgres

server:
	cd server && go run ./cmd/server