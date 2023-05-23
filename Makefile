sqlcli:
	@kubectl exec -it tanzuquiz-0 -- psql

sqlcredentials:
	@./config/db/credentials.sh

start:
	@docker compose up -d
	@echo "login to pgadmin using http://localhost:5050"
	@echo "Test api using http://localhost:8080/api/healthchecker"

stop:
	@docker compose down

clean: stop
	@docker volume rm tanzuquiz_db

build:
	@docker compose build
	@docker compose up -d --force-recreate