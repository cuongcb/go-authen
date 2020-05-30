run:
	@go run ./cmd/go-authen/main.go

base:
	@cd ./deployments && docker-compose up appdb redis &

up:
	@cd ./deployments && docker-compose up go-authen &

down:
	@cd ./deployments && docker-compose stop go-authen

down-base:
	@cd ./deployments && docker-compose stop appdb redis

prepare:
	@cd ./scripts/mysql && ./entry_point.sh
