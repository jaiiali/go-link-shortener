package factory

import (
	"os"

	"github.com/jaiiali/go-link-shortener/internal/core/ports"
	"github.com/jaiiali/go-link-shortener/internal/repositories/link"
)

func NewRepository() ports.LinkRepository {
	var repo ports.LinkRepository

	switch os.Getenv("APP_STORAGE") {
	case "memory":
		repo = link.NewMemRepo()

	case "postgres":
		var err error
		repo, err = link.NewPostgresRepo()
		if err != nil {
			panic(err)
		}

	default:
		panic("storage is not set correctly (memory or postgres)")
	}

	return repo
}
