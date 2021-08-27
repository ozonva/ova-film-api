build:
	go build -o ova-film-api ./cmd/main.go

run:
	./ova-film-api

mockgen:
	mockgen -source=internal/repo/repo.go -destination=./mocks/repo_mock.go -package=mocks github.com/ozonva/ova_film_api/internal/repo Repo

test:
	cd internal/flusher && go test
