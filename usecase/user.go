/*
	Package usecase はドメインを用いて行いたい抽象的な仕事を定義します。
	その仕事をどのように行うかは、interface層以下で実装します。
*/
package usecase

import (
	"fmt"
	"server/domain"
)

type UserRepository interface {
	CreateUser(*domain.User) error
	UpdateUser(*domain.User, int) error
	GetUsers() ([]domain.User, error)
	GetUserByID(int) (domain.User, error)
	DeleteUserByID(int) error
}
type UserUseCase interface {
	CreateUser(*domain.User) error
	UpdateUser(*domain.User, int) error
	GetUsers() ([]domain.User, error)
	GetUserByID(int) (domain.User, error)
	DeleteUserByID(int) error
}
type UserUsecase struct {
	UserRepository UserRepository
}

func (usecase *UserUsecase) CreateUser(u *domain.User) error {
	err := usecase.UserRepository.CreateUser(u)
	if err != nil {
		err = fmt.Errorf("[usecase.CreateUser] failed: %w ", err)
		return err
	}
	return nil
}

func (usecase *UserUsecase) UpdateUser(u *domain.User, id int) error {
	err := usecase.UserRepository.UpdateUser(u, id)
	if err != nil {
		err = fmt.Errorf("[usecase.UpdateUser] failed: %w ", err)
		return err
	}
	return nil
}

func (usecase *UserUsecase) GetUsers() ([]domain.User, error) {
	users, err := usecase.UserRepository.GetUsers()
	if err != nil {
		err = fmt.Errorf("[usecase.GetUsers] failed: %w ", err)
		return nil, err
	}
	return users, nil
}

func (usecase *UserUsecase) GetUserByID(id int) (domain.User, error) {
	user, err := usecase.UserRepository.GetUserByID(id)
	if err != nil {
		err = fmt.Errorf("[usecase.GetUserByID] failed: %w ", err)
		return user, err
	}
	return user, nil
}

func (usecase *UserUsecase) DeleteUserByID(id int) error {
	err := usecase.UserRepository.DeleteUserByID(id)
	if err != nil {
		err = fmt.Errorf("[usecase.DeleteUserByID] failed: %w ", err)
		return err
	}
	return nil
}
