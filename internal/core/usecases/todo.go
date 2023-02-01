package usecases

import (
	"time"

	"github.com/jaiiali/go-link-shortener/helpers"
	"github.com/jaiiali/go-link-shortener/internal/core/domain"
	"github.com/jaiiali/go-link-shortener/internal/core/ports"
	"github.com/jaiiali/go-link-shortener/pkg/logger"
)

type linkUseCase struct {
	repo ports.LinkRepository
	log  *logger.Logger
}

func (t *linkUseCase) FindAll() ([]domain.Link, error) {
	result, err := t.repo.FindAll()
	if err != nil {
		t.log.Error(err)
		return nil, err
	}

	return result, nil
}

func (t *linkUseCase) FindByID(id string) (*domain.Link, error) {
	result, err := t.repo.FindByID(id)
	if err != nil {
		t.log.Error(err)
		return nil, err
	}

	return result, nil
}

func (t *linkUseCase) Create(original string) (*domain.Link, error) {
	link := domain.NewLink(helpers.NewShortID(), original, time.Now())

	result, err := t.repo.Create(link)
	if err != nil {
		t.log.Error(err)
		return nil, err
	}

	return result, nil
}

func NewLinkUseCase(repo ports.LinkRepository, log *logger.Logger) ports.LinkUseCase {
	return &linkUseCase{
		repo: repo,
		log:  log,
	}
}
