package todo

import (
	"errors"
	"time"

	"github.com/oklog/ulid/v2"
)

var (
	ErrInvalidTitle       = errors.New("len title must be greater than 3")
	ErrInvalidDescription = errors.New("len description must be greater than 3")
)

type Todo struct {
	Id          string    `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	Done        bool      `db:"is_done"`
	CreatedAt   time.Time `db:"created_at"`
}

func New() Todo {
	return Todo{}
}

func (t *Todo) LoadFromSaveRequest(title string, description string, createdAt time.Time) {
	t.Id = ulid.Make().String()
	t.Title = title
	t.Description = description
	t.CreatedAt = createdAt
	t.Done = t.IsDone()
}

func (t Todo) IsDone() bool {
	return time.Now().After(t.CreatedAt)
}

func (t Todo) Validate() (err error) {
	// if len(t.Title) <= 3 {
	// 	return ErrInvalidTitle
	// }

	// if len(t.Description) <= 3 {
	// 	return ErrInvalidDescription
	// }
	return
}
