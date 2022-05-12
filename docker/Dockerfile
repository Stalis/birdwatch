############################
# STEP 1 build golang server
############################
FROM golang:alpine as builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

WORKDIR /app
COPY . .

RUN go mod tidy -v
RUN GOOS=linux go build -ldflags="-w -s" -v -o /app/bin/server ./cmd/server

############################
# STEP 2 build a small image
############################
FROM scratch

COPY --from=builder /app/bin/server /bin/server

EXPOSE 50051

ENTRYPOINT [ "/bin/server" ]
CMD [ "--port", "50051", "--logging.file", "/opt/server.log", "--logging.level", "Debug", "--logging.verbose" ]