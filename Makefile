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

deploy-k8s:
	kubectl apply -f k8s/redis/redis.yaml
	kubectl wait --for=condition=available deployment/redis
	kubectl apply -f k8s/broadcast-service/broadcast-service.yaml
	kubectl wait --for=condition=available deployment/broadcast-service
	kubectl apply -f k8s/browser-service/browser-service.yaml
	kubectl wait --for=condition=available deployment/browser-service
	kubectl port-forward service/browser-service 32323:3000 > /dev/null 2>&1 &
	open ./index.html

stop-portforward:
	pkill kubectl -9
