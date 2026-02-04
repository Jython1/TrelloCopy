.PHONY: test-board-mem test-board-race-mem test-card-mem test-card-race-mem test-column-mem test-column-race-mem
.PHONY: test-board-pg test-board-race-pg test-card-pg test-card-race-pg test-column-pg test-column-race-pg
.PHONY: docker-build docker-run docker-stop docker-clean docker-logs
.PHONY: docker-compose-up docker-compose-down docker-compose-logs docker-compose-build

#Testing In-Memory
test-board-mem:
	go test ./internal/repository/in_memory_repository/board_repository_mem_test.go ./internal/repository/in_memory_repository/board_repository_mem.go -v

test-board-race-mem:
	go test ./internal/repository/in_memory_repository/board_repository_mem_test.go ./internal/repository/in_memory_repository/board_repository_mem.go -race -v

test-column-mem:
	go test ./internal/repository/in_memory_repository/column_repository_mem_test.go ./internal/repository/in_memory_repository/column_repository_mem.go -v

test-column-race-mem:
	go test ./internal/repository/in_memory_repository/column_repository_mem_test.go ./internal/repository/in_memory_repository/column_repository_mem.go -v -race

test-card-mem:
	go test ./internal/repository/in_memory_repository/card_repository_mem_test.go ./internal/repository/in_memory_repository/card_repository_mem.go -v

test-card-race-mem:
	go test ./internal/repository/in_memory_repository/card_repository_mem_test.go ./internal/repository/in_memory_repository/card_repository_mem.go -v -race
	
#Testing PostgreSQL
test-board-pg:
	go test ./internal/repository/pg_repository/board_repository_pg_test.go ./internal/repository/pg_repository/board_repository_pg.go -v

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
