swag:
	swag fmt && swag init

up-build:
	docker-compose up --build

up:
	docker-compose up

up-deattached:
	docker-compose up -d

logs:
	docker-compose logs -f

down:
	docker-compose down

