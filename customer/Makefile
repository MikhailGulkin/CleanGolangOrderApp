include .env.dev
up-containers:
	docker-compose -p customer-app -f ./docker-compose.dev.yaml up --build
migrate-new:
	migrate create -ext sql -dir ${MIGRATION_PATH} -seq ${MIGRATION_NAME}
migrate-up:
	migrate -path ${MIGRATION_PATH} -database ${DB_URL} -verbose up
migrate-down:
	migrate -path ${MIGRATION_PATH} -database ${DB_URL} -verbose down
migrate-force:
	migrate -path ${MIGRATION_PATH} -database ${DB_URL} -verbose force ${VERSION}