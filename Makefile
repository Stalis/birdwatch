docker_tags = v0.0.1
server_main = cmd/server/main.go
server_bin = bin/server
covermode = set
coverprofile = coverage.out

proto_src_path = api
proto_dest_path = pkg/api/pb
protoc_opts = --proto_path=$(proto_src_path) --go_out=$(proto_dest_path) --go-grpc_out=$(proto_dest_path)

grpc_port = 50051

.PHONY: init grpc build clean run coverage_html get mod-tidy docker

init: 
	git config core.hooksPath .githooks

grpc:
	protoc $(protoc_opts) $(proto_src_path)/*.proto

build: mod-tidy grpc
	go build -o $(server_bin) $(server_main)

clean:
	rm -rf ./bin $(coverprofile)

run: mod-tidy
	go run $(server_main) --port $(grpc_port) --logging-level Debug --logging-verbose --logging-file ./logs/server.log

test: mod-tidy
	go test -v ./...

coverage_html: $(coverprofile)
	go tool cover -html=$(coverprofile)

$(coverprofile): mod-tidy
	go test -covermode=$(covermode) -coverprofile=$(coverprofile) ./...

mod-tidy:
	go mod tidy

docker:
	docker build ./docker/Dockerfile --tag $(docker_tags)

