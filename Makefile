
run:
	go run main.go

dev: dev-build
	@echo "Starting development server"
	docker-compose up --watch

dev-build:
	@echo "Starting development server"
	docker-compose down
	docker-compose build

build:
	@echo "Building export..."
	go build -o ./build/exort .
	@echo "Complete!"


