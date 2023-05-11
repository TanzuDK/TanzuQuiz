sqlcli:
	@kubectl exec -it tanzuquiz-0 -- psql

sqlcredentials:
	@./config/db/credentials.sh

start:
	@docker compose up -d
	@echo "login to pgadmin using http://localhost:5050"

stop:
	@docker compose down

clean: stop
	@docker volume rm tanzuquiz_db