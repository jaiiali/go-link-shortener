package ports

import (
	"github.com/jaiiali/go-link-shortener/internal/core/domain"
)

type LinkUseCase interface {
	FindAll() ([]domain.Link, error)
	FindByID(id string) (*domain.Link, error)
	Create(original string) (*domain.Link, error)
}
