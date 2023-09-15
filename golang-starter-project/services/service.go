package services

import "golang-starter-project/repositories"

type Service struct {
	repository repositories.Repository
}

func (s *Service) SetDependencies(repository repositories.Repository) {
	s.repository = repository
}

func (s *Service) DoSomething() string {
	return "Some response"
}
