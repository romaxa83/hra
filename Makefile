.SILENT:
.PHONY:
#=================================
# Run service

run:
	@echo Run app
	go run cmd/main.go -config=./config/config.yml

#=================================
# Command for docker

up: up_docker

info: ps info_domen

up_docker:
	docker-compose up

up_docker_d:
	docker-compose up -d

down:
	docker-compose down --remove-orphans

rebuild: down build up_docker info

# флаг -v удаляет все volume (очищает все данные)
down-clear:
	docker-compose down -v --remove-orphans

build:
	docker-compose build

ps:
	docker-compose ps

# ================================
# MongoDB

mongo:
	cd ./scripts && mongo admin -u admin -p admin < mongo_init.js
#=================================
# Proto

proto: proto_order

proto_order:
	@echo Generating order proto
	cd proto && protoc --go_out=. --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=. order.proto