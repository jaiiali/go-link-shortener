package link

import (
	"errors"

	"github.com/jaiiali/go-link-shortener/internal/core/domain"
)

type MemRepo struct {
	data map[string]memDB
}

func NewMemRepo() *MemRepo {
	return &MemRepo{
		data: make(map[string]memDB),
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

	return nil, errors.New("link not found")
}

func (r *MemRepo) FindByOriginal(original string) (*domain.Link, error) {
	for _, link := range r.data {
		if link.Original == original {
			return link.unbind(), nil
		}
	}

	return nil, errors.New("link not found")
}

func (r *MemRepo) Create(link *domain.Link) (*domain.Link, error) {
	var storage = &memDB{}
	storage.bind(link)

	r.data[link.ID] = *storage

	return link, nil
}
