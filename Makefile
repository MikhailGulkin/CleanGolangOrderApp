up-containers:
	docker-compose -f ./docker-compose.dev.yaml up --build
up-container-tests:
	docker-compose -p test-container-app -f ./docker-compose.container.test.yaml up  --build
up-tests:
	docker-compose -p test-app -f ./docker-compose.test.yaml up  --build
up-prod:
	docker-compose -p prod -f ./docker-compose.yaml up --build