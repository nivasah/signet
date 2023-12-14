package bookmark

import (
	"context"

	"github.com/nivasah/signet/internal/resource"
)

type Bookmark struct {
	Name        string
	Location    string
	Description string
	Tags        string
	FolderID    uint
	UserID      uint
}

type Service interface {
	CreateBookmark(b *Bookmark) error
}

func NewService(r resource.BookmarkRepository) Service {
	return &bookmarkSvc{
		Repo: r,
	}
}

type bookmarkSvc struct {
	Repo resource.BookmarkRepository
}

// CreateBookmark implements Service.
func (s *bookmarkSvc) CreateBookmark(b *Bookmark) error {
	bk := resource.Bookmark{
		Name:        b.Name,
		Location:    b.Location,
		Description: b.Description,
		Tags:        b.Tags,
		ID:          0,
		DirID:       0,
	}
	return s.Repo.Create(context.Background(), &bk)
}
