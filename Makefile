docker-run: docker-build
	docker run -p 8000:8000 go-app

docker-build:
	docker build -t go-app .

swag-init:
	swag init -g cmd/main.go

run:
	go run cmd/main.go