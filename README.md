### Showcase setup 
    
    cp .env.dist .env
    docker-compose up
    
    
API is exposed on **127.0.0.1:50000**, use the postman collection at api/postman/market ledger.postman_collection.json to navigate endpoints


### Development

    cp .env.dist .env
    
    docker-compose -f docker-compose.dev.yml up -d
    
    go get -u -d ./...
    
    go mod vendor
    
    go test -json ./...
    
    # expose api
    go run -mod=vendor cmd/ledger/ledger.go serve --config "configs/ledger.dev.yml" --purge-db
    
    # If changes are made to the services
    
    # *nix
    protoc -I . -I third_party api\proto\v1\market-ledger-service.proto --go-grpc_out=. --grpc-gateway_out=logtostderr=true:. --go_out=./
    
    # Windows
    protoc -I . -I third_party api/proto/v1/market-ledger-service.proto --go-grpc_out=. --grpc-gateway_out=logtostderr=true:. --go_out=./
    
    
    
    