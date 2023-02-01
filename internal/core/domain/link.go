package domain

import (
	"fmt"
	"time"
)

type Link struct {
	ID        string
	Original  string
	CreatedAt time.Time
}

func NewLink(id, original string, createdAt time.Time) *Link {
	return &Link{
		ID:        id,
		Original:  original,
		CreatedAt: createdAt,
	}
}

func (t *Link) String() string {
	return fmt.Sprintf("%s: %s", t.ID, t.Original)
}
