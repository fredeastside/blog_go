source .env

MIGRATE_CMD=migrate -source file://./migrations -database $(DATABASE_URL)

run:
	go run cmd/main.go

rundb:
	docker-compose up -d

stopdb:
	docker-compose down

migrate:
	$(MIGRATE_CMD) up

rollback:
	$(MIGRATE_CMD) down

build:
	cd cmd && go build -o ../$(PROJECT_NAME)

build-linux:
	cd cmd && GOOS=linux GOARCH=amd64 go build -o ./$(PROJECT_NAME)

start: build
	./$(PROJECT_NAME)

docker-build:
	docker build -t $(PROJECT_NAME) .

docker-run:
	docker run -p 80:80 -v -m=2G --memory-swap=2G $(PROJECT_NAME)

docker-stats:
	docker stats $(PROJECT_NAME)

pprof: build
	time ./$(PROJECT_NAME) > memory.profile

run-frontend:
	cd frontend && npm run serve

build-frontend:
	cd frontend && npm run build
