# run specific env
go run main.go -GO_ENV prod
go run cmd/order-processing/main.go -GO_ENV dev
go run main.go -GO_ENV test

# test recursively
go test ./...

# show all keys in sorted set
zrangebyscore "sorted set index" -inf +inf

# start redis instance on port
redis-server --port 6385
