package link

import (
	"crypto/md5"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" //nolint
	_ "github.com/lib/pq"                                //nolint

	"github.com/jaiiali/go-link-shortener/internal/core/domain"
	"github.com/jaiiali/go-link-shortener/pkg/logger"
)

type PostgresRepo struct {
	db     *sql.DB
	logger *logger.Logger
}

func NewPostgresRepo() (*PostgresRepo, error) {
	connStr := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_DRIVER"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_SSLMODE"),
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("The connection to the database was successful.")

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		"postgres",
		driver,
	)
	if err != nil {
		return nil, err
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return nil, err
	}

	log.Println("The migration was done successfully.")

	return &PostgresRepo{
		db:     db,
		logger: logger.NewLogger(),
	}, nil
}

func (r *PostgresRepo) Close() error {
	return r.db.Close()
}

func (r *PostgresRepo) FindAll() ([]domain.Link, error) {
	query := `
	SELECT
		id, original, created_at
	FROM links
	ORDER BY created_at DESC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var result = []domain.Link{}

	for rows.Next() {
		var row domain.Link

		err = rows.Scan(&row.ID, &row.Original, &row.CreatedAt)
		if err != nil {
			return nil, err
		}

		result = append(result, row)
	}

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *PostgresRepo) FindByID(id string) (*domain.Link, error) {
	query := `
	SELECT
		id, original, created_at
	FROM links
	WHERE id = $1
	`

	var row = domain.Link{}
	err := r.db.QueryRow(query, id).Scan(&row.ID, &row.Original, &row.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &row, nil
}

func (r *PostgresRepo) FindByOriginal(original string) (*domain.Link, error) {
	query := `
	SELECT
		id, original, created_at
	FROM links
	WHERE hashed = $1
	LIMIT 1
	`

	hashed := fmt.Sprintf("%x", md5.Sum([]byte(original)))

	var row = domain.Link{}
	err := r.db.QueryRow(query, hashed).Scan(&row.ID, &row.Original, &row.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &row, nil
}

func (r *PostgresRepo) Create(link *domain.Link) (*domain.Link, error) {
	query := `
	INSERT INTO links (id, hashed, original, created_at)
	VALUES ($1, $2, $3, $4)
	`

	hashed := fmt.Sprintf("%x", md5.Sum([]byte(link.Original)))

	_, err := r.db.Exec(query, link.ID, hashed, link.Original, link.CreatedAt)
	if err != nil {
		return nil, err
	}

	return link, nil
}
