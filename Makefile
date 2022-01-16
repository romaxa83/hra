.SILENT:
.PHONY:
#=================================
# Run service

run:
	@echo Run app
	go run cmd/main.go -config=./config/config.yml

#=================================
# Proto

proto: proto_order

proto_order:
	@echo Generating order proto
	cd proto && protoc --go_out=. --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=. order.proto