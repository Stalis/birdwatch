.DEFAULT_TARGET := build

docker_version = v0.0.3
docker_flags = -t birdwatch:$(docker_version) -t birdwatch:latest

server_main = cmd/server/main.go
server_bin = bin/server
client_main = cmd/client/main.go
client_bin = bin/client
covermode = set
coverprofile = coverage.out

proto_src_path = api
proto_dest_path = pkg/api/pb
protoc_opts = --proto_path=$(proto_src_path) --go_out=$(proto_dest_path) --go-grpc_out=$(proto_dest_path)

grpc_port = 50051
run_opts = --port $(grpc_port) --logging-level Debug --logging-verbose --logging-file ./logs/server.log
client_opts = --port $(grpc_port) --host localhost

.PHONY: init mod-tidy lint grpc
init: 
	git config core.hooksPath .githooks

mod-tidy:
	go mod tidy

lint:
	golangci-lint run ./...

grpc:
	protoc $(protoc_opts) $(proto_src_path)/*.proto

.PHONY: build clean run test
build: $(server_bin)
	
$(server_bin): mod-tidy grpc
	go build -o $(server_bin) $(server_main)

clean:
	rm -rf ./bin $(coverprofile)

run: mod-tidy
	go run $(server_main) $(run_opts)

test: mod-tidy
	go test -v ./...

.PHONY: coverage_html
coverage_html: $(coverprofile)
	go tool cover -html=$(coverprofile)

$(coverprofile): mod-tidy
	go test -covermode=$(covermode) -coverprofile=$(coverprofile) ./...

.PHONY: docker docker-scratch
docker:
	docker build $(docker_flags) -f ./docker/Dockerfile .

docker-scratch:
	docker build $(docker_flags) -f ./docker/Dockerfile.scratch .

.PHONY: client-run client-build
client_build:
	go build $(client_main) -o $(client_bin)

client-run:
	go run $(client_main) $(client_opts)

.PHONY: integration
integration:
	go run ./integration