package link

import (
	"time"

	"github.com/jaiiali/go-link-shortener/internal/core/domain"
)

type memDB struct {
	ID        string
	Original  string
	CreatedAt time.Time
}

func (l *memDB) bind(link *domain.Link) {
	l.ID = link.ID
	l.Original = link.Original
	l.CreatedAt = link.CreatedAt
}

func (l *memDB) unbind() *domain.Link {
	return &domain.Link{
		ID:        l.ID,
		Original:  l.Original,
		CreatedAt: l.CreatedAt,
	}
}
