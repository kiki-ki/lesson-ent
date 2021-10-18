run_ent:
	go run ./app/ent/cmd/main.go

run_gorm:
	go run ./app/gorm/cmd/main.go

lint:
	golangci-lint run

fmt:
	goimports -w ./
	go vet ./...
	go fmt ./...
	make lint

ent_update:
	go generate ./ent