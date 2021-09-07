package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/ozonva/ova_film_api/internal/repo"
	"github.com/ozonva/ova_film_api/pkg/generated/api"
	"github.com/ozonva/ova_film_api/pkg/ova_film_api"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

const (
	grpcPort           = ":8082"
	grpcServerEndpoint = "localhost:8082"
)

func run(db *sqlx.DB) error {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	api.RegisterMovieServiceServer(s, ova_film_api.NewServer(repo.NewRepo(db)))

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}

func runJSON() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := api.RegisterMovieServiceHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		panic(err)
	}

	err = http.ListenAndServe(":8081", mux)
	if err != nil {
		panic(err)
	}
}

func main() {

	dsn := "postgresql://ovafilm:ovafilm@localhost:5432/ovafilm"
	db, err := sqlx.Open("pgx", dsn) // *sql.DB
	if err != nil {
		log.Fatalf("failed to load driver: %v", err)
	}

	err = db.PingContext(context.Background())
	if err != nil {
		xerrors.Errorf("failed to connect to db: %v", err)
	}

	go runJSON()

	if err := run(db); err != nil {
		log.Fatal(err)
	}
}
