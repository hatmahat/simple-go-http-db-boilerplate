package service

import (
	"simple-http-boilerplate/model"
	"simple-http-boilerplate/repository"
)

type UserService struct {
	UserRepo *repository.UserRepo
}

type UserServiceParam struct {
	UserRepo *repository.UserRepo
}

func NewUserService(param UserServiceParam) *UserService {
	return &UserService{
		UserRepo: param.UserRepo,
	}
}

func (s *UserService) GetUsers() ([]model.User, error) {
	return s.UserRepo.GetAll()
}

func (s *UserService) AddUser(user model.User) (*model.User, error) {
	return s.UserRepo.Create(user)
}
