/*
	Package controller では以下のことを行います

	ビジネスロジックの呼び出し（コントロール）
	- repositoryへDB操作指示。
	- その他

*/
package controller

import (
	"fmt"
	"server/adapter/repository"
	"server/domain"
	"server/ent"
	"server/usecase"
)

type UserController struct {
	Usecase usecase.UserUseCase
}

func NewUserController(db *ent.Client) *UserController {
	return &UserController{
		Usecase: &usecase.UserUsecase{
			UserRepository: &repository.UserRepository{
				DB: db,
			},
		},
	}
}

func (controller *UserController) GetUsers() ([]domain.User, error) {
	users, err := controller.Usecase.GetUsers()
	if err != nil {
		err = fmt.Errorf("[controller.UserController.GetUsers] failed to get users: %w", err)
		return users, err
	}
	return users, err

}

func (controller *UserController) GetUserByID(id int) (domain.User, error) {

	user, err := controller.Usecase.GetUserByID(id)
	if err != nil {
		err = fmt.Errorf("[controller.UserController.GetUserByID] failed to get users: %w", err)
		return user, err
	}
	return user, err

}

func (controller *UserController) CreateUser(user *domain.User) (err error) {

	err = controller.Usecase.CreateUser(user)
	if err != nil {
		err = fmt.Errorf("[controller.UserController.GetUsers] failed to create user: %w", err)
		return err
	}
	return nil
}

func (controller *UserController) UpdateUser(user *domain.User, id int) (err error) {

	err = controller.Usecase.UpdateUser(user, id)
	if err != nil {
		err = fmt.Errorf("[controller.UserController.UpdateUser] failed to update user: %w", err)
		return err
	}
	return nil
}

func (controller *UserController) DeleteUserByID(id int) error {
	err := controller.Usecase.DeleteUserByID(id)
	if err != nil {
		err = fmt.Errorf("[controller.UserController.DeleteUserByID] failed to delete user: %w", err)
		return err
	}
	return nil
}
