#!make
include properties.env
export $(shell sed 's/=.*//' properties.env)

.DEFAULT_GOAL := default

test:
	env

.PHONY: default
default: dependencies build run

.PHONY: dependencies
dependencies:
	govendor list
	govendor add +external
	#govendor fetch +external

.PHONY: build
build: 
	go build -o ${EXECUTABLE} .

.PHONY: docker
docker: 
	docker build -t ${PROJECT}/${PROJECT_NAME}:${VERSION} .
	docker tag ${PROJECT}/${PROJECT_NAME}:${VERSION} ${PROJECT}/${PROJECT_NAME}:latest

.PHONY: dockerhub
dockerhub: 
	docker build -t ${PROJECT}/${PROJECT_NAME}:${VERSION} .
	docker tag ${PROJECT}/${PROJECT_NAME}:${VERSION} ${PROJECT}/${PROJECT_NAME}:latest
	docker push la3mmchen/elastic-cluster-diff:${VERSION}
	docker push la3mmchen/elastic-cluster-diff:latest

run:
	@echo "\n"
	./${EXECUTABLE}

docker-run:
	@echo "\n"
	docker run --rm -it ${PROJECT}/${PROJECT_NAME}:latest --help

compare:
	@echo "\n"
	./${EXECUTABLE} compare --config --cluster localhost:9200 --cluster localhost:8200

