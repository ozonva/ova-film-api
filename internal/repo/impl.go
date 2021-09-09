package repo

import (
	"context"
	"github.com/rs/zerolog/log"

	//"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/ozonva/ova_film_api/internal/movies"
)

func NewRepo(db *sqlx.DB) MovieRepo {
	return &repoImpl{
		db:  db,
		ctx: context.Background(),
	}
}

type repoImpl struct {
	db  *sqlx.DB
	ctx context.Context
}

func (r *repoImpl) ListMovies(limit, offset uint64) ([]movies.Movie, error) {

	var list []movies.Movie

	err := r.db.Select(&list, "select id, user_id, title, year from movies limit $1 offset $2", limit, offset)

	if err != nil {
		log.Printf("list of movies requesting thrown an error %s", err)
		log.Err(err)
	}

	return list, nil
}

func (r *repoImpl) DescribeMovie(entityId uint64) (*movies.Movie, error) {
	var received movies.Movie

	err := r.db.Get(&received, "select id, user_id, title, year from movies where id = $1", entityId)

	if err != nil {
		log.Printf("movie requesting thrown an error %s", err)
		log.Err(err)
	}

	return &received, nil
}

func (*repoImpl) AddMovies(entities []movies.Movie) error {
	return nil
}

func (r *repoImpl) AddMovie(entity *movies.Movie) error {
	_, err := r.db.Exec("insert into movies (user_id, title, year) values ($1, $2, $3)", entity.UserId, entity.Name, entity.Year)
	if err != nil {
		log.Printf("movie adding thrown an error %s", err)
		log.Err(err)
	}
	return nil
}
func (r *repoImpl) RemoveMovie(entityId uint64) error {
	_, err := r.db.Exec("delete from movies where id = $1", entityId)
	if err != nil {
		log.Printf("movie removing thrown an error %s", err)
		log.Err(err)
	}
	return nil
}
