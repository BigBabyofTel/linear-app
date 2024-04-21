dev-db:
	@docker compose -f docker-compose.dev.yaml up -d

run-prod:
	@docker compose -f docker-compose.prod.yaml up -d

down-prod:
	@docker compose -f docker-compose.prod.yaml down
