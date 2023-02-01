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

func (l *linkUseCase) FindAll() ([]domain.Link, error) {
	result, err := l.repo.FindAll()
	if err != nil {
		l.log.Error(err)
		return nil, err
	}

	return result, nil
}

func (l *linkUseCase) FindByID(id string) (*domain.Link, error) {
	result, err := l.repo.FindByID(id)
	if err != nil {
		l.log.Error(err)
		return nil, err
	}

	return result, nil
}

func (l *linkUseCase) Create(original string) (*domain.Link, error) {
	// Check existence
	link, _ := l.repo.FindByOriginal(original)
	if link != nil {
		return link, nil
	}

	// Create new
	link = domain.NewLink(helpers.NewShortID(), original, time.Now())

	result, err := l.repo.Create(link)
	if err != nil {
		l.log.Error(err)
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
