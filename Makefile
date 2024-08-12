build: 
	go build -o study-mongodb ./cmd/study-mongodb/main.go

run:
	@go run cmd/study-mongodb/main.go

up:
	@echo "Starting containers..."
	docker-compose up --build -d

down:
	@echo "Stoping containers..."
	docker-compose down