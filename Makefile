dev:
	go run . -env-file .env.dev
build:
	rm -rf bin && go build -ldflags "-s -w" -o bin/build .
start:
	./bin/build -env-file .env.dev
start-docker:
	docker compose up -d
stop-docker:
	docker compose down
start-docker-withBuild:
	docker compose up --build --force-recreate -d
create-docs:
	swag init