.PHONY: build
build: vendor-proto .generate .build

.PHONY: .generate
.generate:
		rm -rf pkg/generated
		mkdir -p swagger
		mkdir -p pkg/generated
		protoc -I vendor.protogen \
				--go_out=pkg/generated --go_opt=paths=import \
				--go-grpc_out=pkg/generated --go-grpc_opt=paths=import \
				--grpc-gateway_out=pkg/generated \
				--grpc-gateway_opt=logtostderr=true \
				--grpc-gateway_opt=paths=import \
				--validate_out lang=go:pkg/generated \
				--swagger_out=allow_merge=true,merge_file_name=api:swagger \
				--proto_path=pkg/proto/ service.proto
		mv pkg/generated/github.com/ozonva/ova_film_api/* pkg/generated/
		rm -rf pkg/generated/github.com

.PHONY: .build
.build:
		CGO_ENABLED=0 GOOS=linux go build -o bin/ova_film_api cmd/ova_film_api/main.go

.PHONY: install
install: build .install

.PHONY: .install
install:
		go install cmd/grpc-server/main.go

.PHONY: vendor-proto
vendor-proto: .vendor-proto

.PHONY: .vendor-proto
.vendor-proto:
		mkdir -p vendor.protogen
		mkdir -p vendor.protogen/api/service
		yes | cp -rf pkg/proto/service.proto vendor.protogen/api/service
		@if [ ! -d vendor.protogen/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
			mkdir -p  vendor.protogen/google/ &&\
			mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
			rm -rf vendor.protogen/googleapis ;\
		fi
		@if [ ! -d vendor.protogen/github.com/envoyproxy ]; then \
			mkdir -p vendor.protogen/github.com/envoyproxy &&\
			git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/github.com/envoyproxy/protoc-gen-validate ;\
		fi


.PHONY: deps
deps: install-go-deps

.PHONY: install-go-deps
install-go-deps: .install-go-deps

.PHONY: .install-go-deps
.install-go-deps:
		ls go.mod || go mod init
		go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
		go get -u github.com/golang/protobuf/proto
		go get -u github.com/golang/protobuf/protoc-gen-go
		go get -u google.golang.org/grpc
		go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
		go get github.com/envoyproxy/protoc-gen-validate
		go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
		go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
