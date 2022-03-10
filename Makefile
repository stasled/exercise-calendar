include ./config/.env

## db: Run docker container with Postgresql and make migrations.
db: postgres create-db migrate-up

## stop-db: Stop docker container with Postgresql.
stop-db:
	@docker stop $(POSTGRES_IMAGE_NAME)

postgres:
	@docker run -it --rm \
	--name $(POSTGRES_IMAGE_NAME) \
	-p $(POSTGRES_PORT):$(POSTGRES_PORT) \
	-e POSTGRES_USER=$(POSTGRES_USER) \
	-e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) \
	-d postgres:14
	@echo "Wait for the database to start..."
	sleep 3

create-db:
	@docker exec -it $(POSTGRES_IMAGE_NAME) \
	createdb \
	--username=$(POSTGRES_USER) \
	--owner=$(POSTGRES_USER) \
	$(POSTGRES_DB)

migrate-up:
	@migrate -path ./internal/storage/migrations \
	-database "postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:\
	$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable" -verbose up

## drop-db: Drop database.
drop-db:
	@docker exec -it $(POSTGRES_IMAGE_NAME) dropdb $(POSTGRES_DB)

## rest: Run rest api app.
rest:
	@go run ./cmd/rest_event/main.go

## grpc: Run grpc app.
grpc:
	@go run ./cmd/rest_event/main.go

help: Makefile
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'

## doc: Generate swagger documentation
doc:
	@swag init -g cmd/rest_event/main.go -o docs/rest/

.PHONY: db stop postgres create-db drop-db migrate-up rest grpc help doc