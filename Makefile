.PHONY:user
user:
	@echo "Starting user..."
	@go run ./server/cmd/user &

.PHONY:api
api:
	@echo "Starting api..."
	@go run ./server/cmd/api &

.PHONY:game
game:
	@echo "Starting game..."
	@go run ./sever/cmd/game &