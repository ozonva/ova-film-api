package movies

import (
	"strconv"
	"strings"
)

type Movie struct {
	Id     uint64
	UserId uint64
	Name   string
	Year   uint
}

func New(id uint64, userId uint64, name string, year uint) *Movie {
	return &Movie{Id: id, UserId: userId, Name: name, Year: year}
}

func (r *Movie) String() string {
	fields := strings.Join([]string{
		strconv.FormatUint(r.Id, 10),
		strconv.FormatUint(r.UserId, 10),
		r.Name,
		strconv.FormatUint(uint64(r.Year), 10),
	}, ", ")
	return strings.Join([]string{"{", fields, "}"}, "")
}
