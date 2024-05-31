.PHONY: check-service

gen-grpc-file: check-service
	protoc --go_out=. --go-grpc_out=. api/proto/$(SERVICE)/$(SERVICE).proto

gen-model: check-service
	#protoc --go_out=. api/proto/$(SERVICE)/model.proto

check-service:
ifndef SERVICE
	$(error SERVICE is not set. Usage: make start service)
endif