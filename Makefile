build:
	go build -o film ./cmd/main.go

run:
	./film

mockgen:
	mockgen -source=internal/repo/repo.go -destination=./mocks/repo_mock.go -package=mocks github.com/ozonva/ova_film_api/internal/repo Repo

test:
	cd internal/flusher && go test

deps:
	ls go.mod || go mod init github.com/ozonva/ova_film_api
	GOBIN=$(LOCAL_BIN) go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	GOBIN=$(LOCAL_BIN) go get -u github.com/golang/protobuf/proto
	GOBIN=$(LOCAL_BIN) go get -u github.com/golang/protobuf/protoc-gen-go
	GOBIN=$(LOCAL_BIN) go get -u google.golang.org/grpc
	GOBIN=$(LOCAL_BIN) go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	GOBIN=$(LOCAL_BIN) go get google.golang.org/protobuf/reflect/protoreflect@v1.27.1
	GOBIN=$(LOCAL_BIN) go get google.golang.org/protobuf/runtime/protoimpl@v1.27.1
	GOBIN=$(LOCAL_BIN) go get github.com/rs/zerolog/log

grpcgen:
	protoc -I vendor.protogen --go_out=pkg/service --go_opt=paths=import --go-grpc_out=pkg/service --go-grpc_opt=paths=import --proto_path=pkg/proto/ service.proto