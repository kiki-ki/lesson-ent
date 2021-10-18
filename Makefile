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