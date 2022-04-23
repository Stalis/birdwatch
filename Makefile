docker_tags = v0.0.1
server_main = cmd/server/main.go
server_bin = bin/server
config_path = configs/.env
covermode = set
coverprofile = coverage.out

.PHONY: build docker get mod-tidy mod-download

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
