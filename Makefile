.PHONY: test-board test-board-race test-card test-card-race test-column test-column-race
.PHONY: docker-build docker-run docker-stop docker-clean docker-logs
.PHONY: docker-compose-up docker-compose-down docker-compose-logs docker-compose-build

#Testing modeles
test-board:
	go test ./internal/repository/board_repository_mem_test.go ./internal/repository/board_repository_mem.go -v

test-board-race:
	go test ./internal/repository/board_repository_mem_test.go ./internal/repository/board_repository_mem.go -race -v

test-column:
	go test ./internal/repository/column_repository_mem_test.go ./internal/repository/column_repository_mem.go -v

test-column-race:
	go test ./internal/repository/column_repository_mem_test.go ./internal/repository/column_repository_mem.go -v -race

test-card:
	go test ./internal/repository/card_repository_mem_test.go ./internal/repository/card_repository_mem.go -v

test-card-race:
	go test ./internal/repository/card_repository_mem_test.go ./internal/repository/card_repository_mem.go -v -race

#Docker commands
docker-build:
	docker build -t trellocopy .

docker-run:
	docker run -d -p 8080:8080 --name trellocopy-container trellocopy

docker-run-dev:
	docker run -d -p 8080:8080 --name trellocopy-container-dev -v $(PWD):/app trellocopy

docker-stop:
	docker stop trellocopy-container || true

docker-rm:
	docker rm trellocopy-container || true

docker-clean: docker-stop docker-rm
	docker rmi trellocopy || true

docker-logs:
	docker logs trellocopy-container

docker-logs-follow:
	docker logs -f trellocopy-container

docker-shell:
	docker exec -it trellocopy-container /bin/sh

docker-restart: docker-stop docker-rm docker-build docker-run

docker-ps:
	docker ps --format "table {{.ID}}\t{{.Names}}\t{{.Status}}\t{{.Ports}}"

docker-compose-up:
	docker compose up -d

docker-compose-down:
	docker compose down

docker-compose-logs:
	docker compose logs -f

docker-compose-build:
	docker compose build
