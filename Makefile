include .env
docker-build:
	docker-compose up --build
run:
	go run cmd/main.go

bench:
	go test -bench . ./cmd/task_2/
