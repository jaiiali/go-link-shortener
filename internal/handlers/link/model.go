package link

import (
	"fmt"
	"time"

	"github.com/jaiiali/go-link-shortener/internal/core/domain"
)

type JSONTime time.Time

func (t JSONTime) String() string {
	return fmt.Sprintf("%q", time.Time(t).Format("2006-01-02 15:04:05"))
}

func (t JSONTime) MarshalJSON() ([]byte, error) {
	return []byte(t.String()), nil
}

type linkReq struct {
	URL string `json:"url" validate:"required"`
}

type linkResp struct {
	ID        string   `json:"id"`
	Original  string   `json:"original"`
	CreatedAt JSONTime `json:"created_at"`
}

func (l *linkResp) bind(link *domain.Link) {
	l.ID = link.ID
	l.Original = link.Original
	l.CreatedAt = JSONTime(link.CreatedAt)
}

type linkListResp []linkResp

func (l *linkListResp) bind(items []domain.Link) {
	for _, v := range items {
		link := linkResp{}
		link.bind(&v)
		*l = append(*l, link)
	}
}
