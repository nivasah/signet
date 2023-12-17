package resource

import "context"

type Bookmark struct {
	Name        string
	Location    string
	Description string
	Tags        string
	ID          uint
	DirID       uint // Foreign key referencing the ID of the dir_table
}

type BookmarkRepository interface {
	Create(ctx context.Context, b *Bookmark) error
}
