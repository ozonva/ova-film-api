build:
	go build -o ova-film-api ./cmd/main.go

run:
	./ova-film-api

mockgen:
	mockgen -source=internal/repo/repo.go -destination=./mocks/repo_mock.go -package=mocks github.com/ozonva/ova_film_api/internal/repo Repo

test:
	cd internal/flusher && go test

.PHONY: deps
deps: .install-go-deps

.PHONY: .install-go-deps
.install-go-deps:
	ls go.mod || go mod init gitlab.com/siriusfreak/lecture-6-demo
	GOBIN=$(LOCAL_BIN) go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	GOBIN=$(LOCAL_BIN) go get -u github.com/golang/protobuf/proto
	GOBIN=$(LOCAL_BIN) go get -u github.com/golang/protobuf/protoc-gen-go
	GOBIN=$(LOCAL_BIN) go get -u google.golang.org/grpc
	GOBIN=$(LOCAL_BIN) go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

grpcgen:
	protoc -I vendor.protogen --go_out=pkg/service --go_opt=paths=import --go-grpc_out=pkg/service --go-grpc_opt=paths=import --proto_path=pkg/proto/ service.proto