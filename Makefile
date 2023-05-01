BROKER_BINARY=brokerService

up:
		@echo "Starting Docker images...."
		docker compose up -d
		@echo "All containers started...."

up_build:
		@echo "Stopping docker images (If Running...)"
		docker compose down
		@echo "Build docker images and run...."
		docker compose up -d --build
		@echo "Success in running latest images..."

down:
		@echo "Stopping docker images...."
		docker compose down
		@echo "Done!"

build_broker:
		@echo "Building standalone broker service..."
		cd ./broker && env GOOS=darwin CGO_ENABLED=0 GOARCH=arm64 go build -o ${BROKER_BINARY} ./cmd/api
		@echo "Done!!"

start_broker: build_broker
		@echo "Starting backend"
		cd ./broker && ./${BROKER_BINARY} &