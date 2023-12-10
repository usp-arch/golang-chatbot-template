APP_NAME=my-app

run:
	go run cmd/${APP_NAME}/main.go

apply-secrets:
	kubectl apply -f secret.yaml

deploy:
	kubectl apply -f deployment.yaml

login:
	docker login registry.my_registry.ru:5000 -u USERNAME -p PASSWORD