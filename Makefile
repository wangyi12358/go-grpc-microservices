include .env
export

.PHONY: check-service

gen-grpc-file: check-service
	protoc --go_out=. --go-grpc_out=. api/proto/$(SERVICE)/$(SERVICE).proto

gen-model:
	go run ./script/gen_db_model.go

start: check-service
	go run cmd/$(SERVICE)/main.go

ent-generate:
	go run -mod=mod cmd/ent/generate/main.go

gql-generate:
	go run -mod=mod github.com/99designs/gqlgen


check-service:
ifndef SERVICE
	$(error SERVICE is not set. Usage: make start service)
endif