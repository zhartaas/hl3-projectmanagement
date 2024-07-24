package management

import (
	"github.com/google/uuid"
	"hl3-projectmanagement/internal/domain/user"
)

func (s *Service) CreateUser(req user.Request) (err error) {
	data := user.Entity{
		ID:    uuid.New().String(),
		Name:  req.Name,
		Email: req.Email,
		Role:  req.Role,
	}

	err = s.UserRepository.Create(data)

	return nil
}

func (s *Service) GetUsers() (res []user.Entity, err error) {
	res, err = s.UserRepository.GetAll()
	return
}

func (s *Service) GetUserByID(id string) (res user.Entity, err error) {
	res, err = s.UserRepository.GetByID(id)

	return
}

func (s *Service) UpdateUser(id string, req user.Request) (err error) {
	err = s.UserRepository.Update(id, req)

	return
}

func (s *Service) DeleteUser(id string) (err error) {
	err = s.UserRepository.Delete(id)

	return
}

func (s *Service) SearchUser(param string, searchByName bool) (user user.Entity, err error) {
	user, err = s.UserRepository.Search(param, searchByName)

	return
}

func (s *Service) UserExists(id string) bool {
	_, err := s.UserRepository.GetByID(id)
	if err != nil {
		return false
	}
	return true
}
