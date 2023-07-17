package todo_test

import (
	"fmt"
	"os"
	"sesi-11/domain/todo"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

var (
	service = todo.Service{}
	db      *sqlx.DB
)

func connectDb() (*sqlx.DB, error) {
	dbConn, err := sqlx.Connect("postgres", fmt.Sprintf(
		"host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	))
	if err != nil {
		return nil, err
	}

	if err := dbConn.Ping(); err != nil {
		return nil, err
	}

	_, err = dbConn.Exec(migrateDB)
	if err != nil {
		return nil, err
	}

	db = dbConn
	return db, nil
}

var migrateDB = `
	CREATE TABLE IF NOT EXISTS todos (
		id varchar(16) PRIMARY KEY,
		title varchar(20),
		description varchar(50),
		is_done boolean,
		created_at timestamptz
	);
`

func init() {
	dbConn, err := connectDb()
	if err != nil {
		panic(err)
	}
	db = dbConn

	// repo := todo.NewRepository(db)
	repoMock := NewRepositoryMock()
	service = todo.NewService(repoMock)
}

func TestConnectDB(t *testing.T) {
	require.NotNil(t, db)
}

func TestSave(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		funcSave = func(item todo.Todo) (err error) {
			return
		}

		item := todo.New()
		item.LoadFromSaveRequest("title", "description", time.Now())
		err := service.Save(item)
		require.Nil(t, err)
	})

	t.Run("fail | len title", func(t *testing.T) {
		funcSave = func(item todo.Todo) (err error) {

			return todo.ErrInvalidTitle
		}

		item := todo.New()
		item.LoadFromSaveRequest("tit", "description", time.Now())
		err := service.Save(item)
		require.NotNil(t, err)
		require.Equal(t, err, todo.ErrInvalidTitle)
	})

	t.Run("fail | len description", func(t *testing.T) {
		funcSave = func(item todo.Todo) (err error) {
			return todo.ErrInvalidDescription
		}

		item := todo.New()
		item.LoadFromSaveRequest("tilet", "des", time.Now())
		err := service.Save(item)
		require.NotNil(t, err)
		require.Equal(t, err, todo.ErrInvalidDescription)
	})
}
