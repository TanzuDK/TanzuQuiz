dbname=$(kubectl get secret tanzuquiz-db-secret -o go-template='{{.data.dbname | base64decode}}')
username=$(kubectl get secret tanzuquiz-db-secret -o go-template='{{.data.username | base64decode}}')
password=$(kubectl get secret tanzuquiz-db-secret -o go-template='{{.data.password | base64decode}}')
echo $dbname
echo $username
echo $password