package todo_test

import (
	"log"
	"sesi-11/domain/todo"
)

type RepositoryMock struct{}

var (
	funcSave    func(item todo.Todo) (err error)
	funcUpdate  func(item todo.Todo) (err error)
	funcGetData func() (item todo.Todo, err error)
)

func NewRepositoryMock() RepositoryMock {
	return RepositoryMock{}
}

func (r RepositoryMock) Save(item todo.Todo) (err error) {
	log.Println("using mock")
	return funcSave(item)
}
func (r RepositoryMock) Update(item todo.Todo) (err error) {
	log.Println("using mock")
	return funcUpdate(item)
}

func (r RepositoryMock) GetData() (item todo.Todo, err error) {
	return funcGetData()
}
