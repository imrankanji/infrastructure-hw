build-broadcast:
	docker build --no-cache -t broadcast-service:latest ./broadcast-service

build-browser:
	docker build --no-cache -t browser-service:latest ./browser-service

build-redis:
	docker build -t redis:latest ./redis

up:
	docker compose up

down:
	docker compose down