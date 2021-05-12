package task

type Service interface {
	Index() ([]Task, error)
	Store(input InputTask) (Task, error)
	Show(id InputTaskDetail) (Task, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Index() ([]Task, error) {
	tasks, err := s.repository.SelectAll()
	if err != nil {
		return tasks, err
	}
	return tasks, nil
}

func (s *service) Store(input InputTask) (Task, error) {
	task := Task{}
	task.Name = input.Name
	task.Description = input.Description

	newTask, err := s.repository.Insert(task)
	if err != nil {
		return newTask, err
	}
	return newTask, nil
}

func (s *service) Show(id InputTaskDetail) (Task, error) {
	task, err := s.repository.SelectById(id.ID)
	if err != nil {
		return task, err
	}
	return task, nil
}
