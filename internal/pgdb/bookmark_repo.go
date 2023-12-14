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

func init() {
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUserName := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	dbURL = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUserName, dbPassword, dbHost, dbPort, dbName)
}

func NewBookmarkRepository() (resource.BookmarkRepository, error) {
	var err error
	conn, err = pgx.Connect(context.Background(), dbURL)
	if err != nil {
		err = fmt.Errorf("unable to connect to the DB: %w", err)
		return nil, err
	}
	b := &bookmarkRepository{
		Conn: conn,
	}
	return b, nil
}

type bookmarkRepository struct {
	Conn *pgx.Conn
}

// Create implements resource.BookmarkRepository.
func (r *bookmarkRepository) Create(ctx context.Context, b *resource.Bookmark) error {
	sqlStr := "insert into bookmarks (name, location, description, tags, dir_id) values ($1, $2, $3, $4, $5)"
	cmd, err := r.Conn.Exec(ctx, sqlStr, b.Name, b.Location, b.Description, b.Tags, b.DirID)
	if err != nil {
		err = fmt.Errorf("unable to insert the record: %w", err)
		return err
	}
	log.Printf("successfully inserted %d records", cmd.RowsAffected())
	return nil
}
