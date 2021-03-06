GO_VERSION_SHORT:=$(shell echo `go version` | sed -E 's/.* go(.*) .*/\1/g')
ifneq ("1.17","$(shell printf "$(GO_VERSION_SHORT)\n1.17" | sort -V | head -1)")
$(warning NEED GO VERSION >= 1.17. Found: $(GO_VERSION_SHORT))
endif

GITHUB_PATH=github.com/arttet/rock-paper-scissors-lizard-spock

###############################################################################

.PHONY: all
all: .reqs deps gen build

.PHONY: reqs
reqs: .reqs

.PHONY: deps
deps: .deps-go

.PHONY: gen
gen: .generate-go

.PHONY: build
build:  .build

.PHONY: debug
debug:
	docker stop game-service || true
	docker run --rm -d --publish 40000:40000 --publish 8080:8080 --security-opt=seccomp:unconfined --name game-service game-service:debug

.PHONY: test
test:
	go test -v -timeout 30s -coverprofile cover.out ./...
	go tool cover -func cover.out | grep -v -E '100.0%|total' || echo "OK"
	go tool cover -func cover.out | grep total | awk '{print ($$3)}'

.PHONY: lint
lint:
	buf lint
	golangci-lint run ./...

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: fmt
fmt:
	find . -iname *.go | xargs gofmt -w
	find . -iname *.proto | xargs clang-format -i

.PHONY: cover
cover:
	go tool cover -html cover.out

.PHONY: grpcui
grpcui:
	grpcui -plaintext localhost:8082

.PHONY: image
image: .image

.PHONY: debug-image
debug-image: .debug-image

.PHONY: clean
clean:
	rm -rd ./bin/ || true
	docker container rm -f game-service || true
	docker volume rm $(docker volume ls-q) || true

################################################################################

.PHONY: install
install:
	cd website && npm install && cd ..

.PHONY: start
start:
	npm start --prefix website

.PHONY: deploy
deploy:
	npm run build --prefix website

################################################################################

.PHONY: .deps-go
.deps-go:
	go mod download
	go install github.com/golang/mock/mockgen@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/envoyproxy/protoc-gen-validate@latest
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest

################################################################################

# https://github.com/bufbuild/buf/releases
BUF_VERSION=v1.0.0-rc10

OS_NAME=$(shell uname -s)
OS_ARCH=$(shell uname -m)
GO_BIN=$(shell go env GOPATH)/bin
BUF_EXE=$(GO_BIN)/buf$(shell go env GOEXE)

ifeq ("NT", "$(findstring NT,$(OS_NAME))")
OS_NAME=Windows
endif

.reqs:
	@command -v buf 2>&1 > /dev/null || (echo "Install buf" && \
		mkdir -p "$(GO_BIN)" && \
		curl -k -sSL0 https://github.com/bufbuild/buf/releases/download/$(BUF_VERSION)/buf-$(OS_NAME)-$(OS_ARCH)$(shell go env GOEXE) -o "$(BUF_EXE)" && \
		chmod +x "$(BUF_EXE)")

################################################################################

.generate-go: .generate-mock .generate-game-service

.generate-mock:
	go generate ./...

.generate-game-service: $(eval SERVICE_NAME := game-service) .generate-template

.generate-template:
	@ echo $(SERVICE_NAME)
	@ $(BUF_EXE) generate
	@ cp -R pkg/$(GITHUB_PATH)/pkg/* pkg/
	@ rm -rf pkg/github.com/
	@ cd pkg/$(SERVICE_NAME) && ls go.mod || (go mod init $(GITHUB_PATH)/pkg/$(SERVICE_NAME) && go mod tidy)

################################################################################

.build: .build-game-service

.build-game-service: \
	$(eval SERVICE_NAME := game-service) \
	$(eval SERVICE_MAIN := cmd/$(SERVICE_NAME)/main.go) \
	$(eval SERVICE_EXE  := ./bin/$(SERVICE_NAME)) \
	.build-template

.build-template:
	CGO_ENABLED=0 go build \
		-mod=mod \
		-ldflags=" \
			-X '$(GITHUB_PATH)/internal/config.version=$(VERSION)' \
			-X '$(GITHUB_PATH)/internal/config.commitHash=$(COMMIT_HASH)' \
		" \
		-o $(SERVICE_EXE)$(shell go env GOEXE) $(SERVICE_MAIN)

################################################################################

.image: .image-game-service

.image-game-service: \
	$(eval SERVICE_NAME := game-service) \
	.image-template

.image-template:
	docker build . --file deploy/docker/$(SERVICE_NAME)/Dockerfile --tag $(SERVICE_NAME):dev

################################################################################

.debug-image: .debug-image-game-service

.debug-image-game-service: \
	$(eval SERVICE_NAME := game-service) \
	.debug-image-template

.debug-image-template:
	docker build . --file deploy/docker/$(SERVICE_NAME)/Dockerfile.debug --tag $(SERVICE_NAME):debug
