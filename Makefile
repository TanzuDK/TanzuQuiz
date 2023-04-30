sqlcli:
	kubectl exec -it tanzuquiz-0 -- psql

sqlcredentials:
	./config/db/credentials.sh
