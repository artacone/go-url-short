
.PHONY: run
run:
	docker-compose up --build shortener db


.PHONY: test
test:
	go test -v ./...


.PHONY: lint
lint:
	golangci-lint run ./...


.PHONY: proto
proto:
	protoc  --go_out=internal --go_opt=paths=source_relative \
		--go-grpc_out=internal --go-grpc_opt=paths=source_relative \
		api/api.proto
