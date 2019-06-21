package store

import (
	"time"

	"{{.ModName}}"
	"{{.ModName}}/uuid"
)

type imdbUser struct {
	Mem map[string]{{.ProjName}}.User
}

func newImdbUser() imdbUser {
	m := make(map[string]{{.ProjName}}.User)
	id := uuid.Gen()
	m[id] = {{.ProjName}}.User {
		ID: id,
		Username: "username",
		Fullname: "fullname",
		CreatedAt: time.Now(),
		LastChangedAt: time.Now(),
	}

	s := imdbUser {
		Mem: m,
	}

	return s
}

func (s *imdbUser) GetAll() ([]{{.ProjName}}.User, error) {
	v := make([]{{.ProjName}}.User, 0, len(s.Mem))
	for  _, value := range s.Mem {
		v = append(v, value)
	}

	return v, nil
}