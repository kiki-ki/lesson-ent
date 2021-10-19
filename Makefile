run:
	go run ./main.go

lint:
	golangci-lint run

fmt:
	goimports -w ./
	go vet ./...
	go fmt ./...
	make lint

ent_update:
	go generate ./ent

company_index_users:
	curl localhost:8080/companies/1

company_show:
	curl localhost:8080/companies/1

company_update:
	curl localhost:8080/companies/1 -X PUT -H "Content-Type: application/json" -d '{"name":"chan2"}'

company_delete:
	curl localhost:8080/companies/1 -X DELETE

company_create_with_user:
	curl localhost:8080/companies -X POST -H "Content-Type: application/json" -d '{"companyName":"hoge","userName":"chan", "userEmail":"chan@exa.com", "userComment":"hoge"}'
