run:
	go run ./cmd/api
up:
	docker-compose -f deployments/docker-compose.dev.yml up -d
down:
	docker-compose -f deployments/docker-compose.dev.yml down
