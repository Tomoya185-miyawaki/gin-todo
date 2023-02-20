.PHONY: up
up:
	docker-compose up

.PHONY: build
build:
	docker-compose build

.PHONY: down
down:
	docker-compose down

.PHONY: api
api:
	docker-compose exec api bash

.PHONY: doc
doc:
	docker-compose exec api godoc -http=:9000

.PHONY: front
front:
	docker-compose exec front sh