# run specific env
go run main.go -GO_ENV prod
go run cmd/order-processing/main.go -GO_ENV dev
go run main.go -GO_ENV test

# test recursively
go test ./...
