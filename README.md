# go-grpc-playground

Just playing around with golang and gRPC

### Getting started
Generate go pb files
```bash
./generate.sh
```
### Run server
```bash
# Greet server
go run greet/greet_server/server.go

# Calculator server
go run calculator/calculator_server/server.go
```

### Run client
```bash
# Greet client
go run greet/greet_client/client.go

# Calculator client
go run calculator/calculator_client/client.go
```

- More to come
