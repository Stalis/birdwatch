docker_tags = v0.0.1
server_main = cmd/server/main.go
server_bin = bin/server
config_path = configs/.env
covermode = set
coverprofile = coverage.out

proto_src_path = api
proto_dest_path = pkg/api/pb
protoc_opts = --proto_path=$(proto_src_path) --go_out=$(proto_dest_path) --go-grpc_out=$(proto_dest_path)

.PHONY: init grpc build docker get mod-tidy mod-download

init: 
	git config core.hooksPath .githooks

build: get
	go build -o $(server_bin) $(server_main)

run: get
	go run $(server_main) --config $(config_path)

test: get
	go test -v ./...

coverage_html: $(coverprofile)
	go tool cover -html=$(coverprofile)

$(coverprofile): get
	go test -covermode=$(covermode) -coverprofile=$(coverprofile) ./...

get: mod-tidy

mod-tidy: mod-download
	go mod tidy

mod-download:
	go mod download -x

docker:
	docker build ./docker/Dockerfile --tag $(docker_tags)

grpc:
	protoc $(protoc_opts) $(proto_src_path)/*.proto