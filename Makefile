
build:
	docker build -t catchbook-backend:last -f services/backend/Dockerfile .
	docker build -t catchbook-frontend:last -f services/frontend/Dockerfile .

up:
	docker compose up -d

down:
	docker compose down
