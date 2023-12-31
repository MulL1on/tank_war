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
	@go run ./server/cmd/game &


.PHONY:data
data:
	@echo "Starting data..."
	@go run ./server/cmd/data &