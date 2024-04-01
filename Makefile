YOUSHARE_DSN := $(shell echo $$YOUSHARE_DSN)
run:
	go run ./cmd/api
up:
	docker-compose -f deployments/docker-compose.dev.yml up -d
down:
	docker-compose -f deployments/docker-compose.dev.yml down
migrate-up:
	migrate -path ./migrations -database=$(YOUSHARE_DSN) -verbose up
migrate-down:
	migrate -path ./migrations -database=$(YOUSHARE_DSN) -verbose down
migrate-force:
	migrate -path ./migrations -database=$(YOUSHARE_DSN) -verbose force $(version)
