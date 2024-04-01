build-broadcast:
	docker build --no-cache -t ikanji/broadcast-service:latest ./broadcast-service

build-browser:
	docker build --no-cache -t ikanji/browser-service:latest ./browser-service

build-redis:
	docker build -t redis:latest ./redis

build-nginx:
	docker build --no-cache -t ikanji/browser-service-nginx:latest ./nginx

up:
	docker compose up

down:
	docker compose down