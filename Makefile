sqlcli:
	kubectl exec -it tanzuquiz-0 -- psql

sqlcredentials:
	./config/db/credentials.sh

start:
	docker compose up -d

stop:
	docker compose down

clean: stop
	docker volume rm tanzuquiz_db