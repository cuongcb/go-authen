run:
	@go run ./cmd/go-authen/main.go

base:
	@docker-compose -f deployments/docker-compose.yml up appdb redis &

up:
	@docker-compose -f deployments/docker-compose.yml up go-authen &

down:
	@docker-compose -f deployments/docker-compose.yml stop go-authen

down-base:
	@docker-compose -f deployments/docker-compose.yml stop appdb redis

prepare:
	@cd ./scripts/mysql && ./entry_point.sh
