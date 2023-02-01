package link

import (
	"errors"

	"github.com/jaiiali/go-link-shortener/internal/core/domain"
	"github.com/jaiiali/go-link-shortener/pkg/logger"
)

type MemRepo struct {
	data map[string]memDB
	log  *logger.Logger
}

func NewMemRepo(log *logger.Logger) *MemRepo {
	return &MemRepo{
		data: make(map[string]memDB),
		log:  log,
	}
}

func (r *MemRepo) FindAll() ([]domain.Link, error) {
	var result = []domain.Link{}
	for _, link := range r.data {
		result = append(result, *link.unbind())
	}

	return result, nil
}

func (r *MemRepo) FindByID(id string) (*domain.Link, error) {
	if result, ok := r.data[id]; ok {
		return result.unbind(), nil
	}

	err := errors.New("link not found")
	r.log.Error(err)

	return nil, err
}

func (r *MemRepo) FindByOriginal(original string) (*domain.Link, error) {
	for _, link := range r.data {
		if link.Original == original {
			return link.unbind(), nil
		}
	}

	err := errors.New("link not found")
	r.log.Error(err)

	return nil, err
}

func (r *MemRepo) Create(link *domain.Link) (*domain.Link, error) {
	var storage = &memDB{}
	storage.bind(link)

	r.data[link.ID] = *storage

	return link, nil
}
