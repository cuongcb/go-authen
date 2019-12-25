run:
	@go run ./cmd/go-authen/main.go

up:
	@cd ./deployments && sudo docker-compose up