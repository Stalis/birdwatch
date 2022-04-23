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
RUN go build -v -o /app/bin/server ./cmd/server

############################
# STEP 2 build a small image
############################
FROM scratch

WORKDIR /app
COPY --from=builder /app/bin .

EXPOSE 80

ENTRYPOINT [ "/app/server" ]