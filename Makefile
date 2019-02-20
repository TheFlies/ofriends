PROJECT_NAME=ofriends
BUILD_VERSION=$(shell cat VERSION)
DOCKER_IMAGE=$(PROJECT_NAME):$(BUILD_VERSION)
GO_BUILD_ENV=CGO_ENABLED=0 GOOS=linux GOARCH=amd64
GO_FILES=$(shell go list ./... | grep -v /vendor/)

.SILENT:

all: fmt vet install test

build: build-web
	$(GO_BUILD_ENV) go build -v -o $(PROJECT_NAME)-$(BUILD_VERSION).bin .
install:
	$(GO_BUILD_ENV) go install
vet:
	go vet $(GO_FILES)
fmt:
	go fmt $(GO_FILES)
test:
	go test $(GO_FILES) -cover
integration_test:
	go test -tags=integration $(GO_FILES)
compose: docker
	cd deployment/docker && docker-compose up
docker: vet test build
	mv $(PROJECT_NAME)-$(BUILD_VERSION).bin deployment/docker/$(PROJECT_NAME).bin; \
	mkdir deployment/docker/web
	cp -R web/dist deployment/docker/web/dist
	cp web/index.html deployment/docker/web/index.html
	cp -R web/assets deployment/docker/web/assets
	cd deployment/docker; \
	docker build -t $(DOCKER_IMAGE) .; \
	rm -rf $(PROJECT_NAME).bin 2> /dev/null;\
	rm -rf web 2> /dev/null;
docker_run:
	docker run -p 8080:8080 $(DOCKER_IMAGE)
build-web:
	cd web; \
	npm install; \
	npm run build
build-web-quick:
	cd web; \
	npm run build
run: build-web-quick
	go run main.go