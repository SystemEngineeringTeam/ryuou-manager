.PHONY: run
COMPOSE=docker-compose
DOCKER=docker
DB=ryuou-manager-db
DB_URI="mysql://root:my-secret-pw@tcp(127.0.0.1:33061)/ryuou_manager_db"
MIGRATE_FILES=migrations
v=

run:
	cd backend && go run main.go

run/air:
	air

build:
	cd go && go build -o ../tmp/main main.go

db/up:
	$(COMPOSE) up -d --build

db/build:
	$(COMPOSE) build

db/sh:
	$(DOCKER) exec -it $(DB) bash

db/restart:
	$(COMPOSE) restart $(DB)

db/down:
	$(COMPOSE) rm -fsv $(DB)

logs:
	$(COMPOSE) logs

ps:
	$(COMPOSE) ps


down:
	$(COMPOSE) down -v

migrate/up:
	migrate -path ${MIGRATE_FILES} -database ${DB_URI} up ${v}

migrate/down:
	migrate -path ${MIGRATE_FILES} -database ${DB_URI} down ${v}

migrate/force:
	migrate -path ${MIGRATE_FILES} -database ${DB_URI} force ${v}

migrate/updown: migrate/down migrate/up