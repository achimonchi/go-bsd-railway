package todo

type RepositoryInterface interface {
	Save(item Todo) (err error)
}

type Service struct {
	repository RepositoryInterface
}

func NewService(repo RepositoryInterface) Service {
	return Service{
		repository: repo,
	}
}

func (s Service) Save(item Todo) (err error) {
	if err = item.Validate(); err != nil {
		return
	}
	err = s.repository.Save(item)
	return
}
