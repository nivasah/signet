package pgdb

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/nivasah/signet/internal/resource"
)

var conn *pgx.Conn

var dbURL string

func NewBookmarkRepository() (*BookmarkPGRepository, error) {
	dbURL = os.Getenv("DB_URL")
	var err error
	conn, err = pgx.Connect(context.Background(), dbURL)
	if err != nil {
		err = fmt.Errorf("unable to connect to the DB: %w", err)
		return nil, err
	}
	b := &BookmarkPGRepository{
		Conn: conn,
	}
	return b, nil
}

type BookmarkPGRepository struct {
	Conn *pgx.Conn
}

// Create implements resource.BookmarkRepository.
func (r *BookmarkPGRepository) Create(ctx context.Context, b *resource.Bookmark) error {
	sqlStr := "insert into resources (name, location, description, tags, dir_id) values ($1, $2, $3, $4, $5)"
	cmd, err := r.Conn.Exec(ctx, sqlStr, b.Name, b.Location, b.Description, b.Tags, b.DirID)
	if err != nil {
		err = fmt.Errorf("unable to insert the record: %w", err)
		log.Printf("unable to insert the record: %v", err)
		return err
	}
	log.Printf("successfully inserted %d records", cmd.RowsAffected())
	return nil
}
