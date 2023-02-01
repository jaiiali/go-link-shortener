package ports

import (
	"github.com/jaiiali/go-link-shortener/internal/core/domain"
)

type LinkRepository interface {
	FindAll() ([]domain.Link, error)
	FindByID(id string) (*domain.Link, error)
	FindByOriginal(original string) (*domain.Link, error)
	Create(link *domain.Link) (*domain.Link, error)
}
