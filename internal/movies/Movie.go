package movies

import (
	"strconv"
	"strings"
)

type Movie struct {
	Id     uint64
	UserId uint64 `db:"user_id"`
	Name   string `db:"title"`
	Year   string
}

func New(id uint64, userId uint64, name string, year string) *Movie {
	return &Movie{Id: id, UserId: userId, Name: name, Year: year}
}

func (r *Movie) String() string {
	fields := strings.Join([]string{
		strconv.FormatUint(r.Id, 10),
		strconv.FormatUint(r.UserId, 10),
		r.Name,
		r.Year,
	}, ", ")
	return strings.Join([]string{"{", fields, "}"}, "")
}
