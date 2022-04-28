docker_tags = v0.0.1
server_main = cmd/server/main.go
server_bin = bin/server
covermode = set
coverprofile = coverage.out

proto_src_path = api
proto_dest_path = pkg/api/pb
protoc_opts = --proto_path=$(proto_src_path) --go_out=$(proto_dest_path) --go-grpc_out=$(proto_dest_path)

.PHONY: init grpc build clean run coverage_html get mod-tidy docker

init: 
	git config core.hooksPath .githooks

grpc:
	protoc $(protoc_opts) $(proto_src_path)/*.proto

build: get grpc
	go build -o $(server_bin) $(server_main)

clean:
	rm -rf ./bin $(coverprofile)

run: get
	go run $(server_main)

test: get
	go test -v ./...

coverage_html: $(coverprofile)
	go tool cover -html=$(coverprofile)

$(coverprofile): get
	go test -covermode=$(covermode) -coverprofile=$(coverprofile) ./...

get: mod-tidy

mod-tidy:
	go mod tidy

docker:
	docker build ./docker/Dockerfile --tag $(docker_tags)

