package todo

import (
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return Repository{db: db}
}

func (r Repository) Save(item Todo) (err error) {
	query := `
		INSERT INTO todos (
			id, title, description, is_done, created_at
		) VALUES (
			:id, :title, :description, :is_done, :created_at
		)
	`

	_, err = r.db.NamedExec(query, item)

	return
}
