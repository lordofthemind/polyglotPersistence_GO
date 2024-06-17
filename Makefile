# Environment variables
MG_CONTAINER_NAME ?= polyglotPersistence_MG
PG_CONTAINER_NAME ?= polyglotPersistence_PG
MG_IMAGE_TAG ?= latest
PG_IMAGE_TAG ?= 16-alpine
MG_DB_NAME ?= tradingview
PG_DB_NAME ?= tradingview

# MongoDB commands
crtmg:
	docker run --name $(MG_CONTAINER_NAME) -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=root -e MONGO_INITDB_ROOT_PASSWORD=tradingViewSecret -d mongo:$(MG_IMAGE_TAG)

crtmgdb:
	docker exec -it $(MG_CONTAINER_NAME) mongo -u root -p tradingViewSecret --authenticationDatabase admin --eval "db.getSiblingDB('$(MG_DB_NAME)').createUser({user: 'root', pwd: 'tradingViewSecret', roles: ['readWrite']})"

strmg:
	docker start $(MG_CONTAINER_NAME)

stpmg:
	docker stop $(MG_CONTAINER_NAME)

rmvmg:
	docker rm $(MG_CONTAINER_NAME)

# PostgreSQL commands
crtpg:
	docker run --name $(PG_CONTAINER_NAME) -p 5432:5432 -e POSTGRES_PASSWORD=tradingViewSecret -d postgres:$(PG_IMAGE_TAG)

crtpgdb:
	docker exec -it $(PG_CONTAINER_NAME) createdb -U postgres $(PG_DB_NAME)

strpg:
	docker start $(PG_CONTAINER_NAME)

stppg:
	docker stop $(PG_CONTAINER_NAME)

rmvpg:
	docker rm $(PG_CONTAINER_NAME)


.PHONY: crtmg crtmgdb crtpg crtpgdb strmg stpmg rmvmg strpg stppg rmvpg