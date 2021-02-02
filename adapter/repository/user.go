/*
	Package repository はデータベースへ送る命令を記述します。
	Doman層で記述したRepositoryのinterfaceをここで実装します。
*/
package repository

import (
	"context"
	"database/sql"
	"fmt"
	"server/domain"
	"server/ent"
)

type UserRepository struct {
	DB *ent.Client
}

func (repo *UserRepository) GetUsers() (users []domain.User, err error) {
	ctx := context.Background()
	getUsersRow, err := repo.DB.User.Query().All(ctx)
	if err != nil {
		err = fmt.Errorf("[repository.UserRepository.GetUsers] failed to find all users: %w", err)
		return nil, err
	}
	if len(getUsersRow) == 0 {
		err = fmt.Errorf("[repository.UserRepository.GetUsers] record not found: %w", sql.ErrNoRows)
		return nil, err
	}
	for _, item := range getUsersRow {
		user := domain.User{
			ID:          item.ID,
			FirstName:   item.FirstName,
			LastName:    item.LastName,
			Email:       item.Email,
			PhoneNumber: item.PhoneNumber,
		}
		users = append(users, user)
	}
	return users, nil
}

func (repo *UserRepository) GetUserByID(id int) (user domain.User, err error) {
	ctx := context.Background()
	getUserByIDRow, err := repo.DB.User.Get(ctx, id)
	if err != nil {
		err = fmt.Errorf("[repository.UserRepository.GetUserByID] failed to find user by id=%d: %w ", id, err)
		return user, err
	}
	user = domain.User{
		ID:          getUserByIDRow.ID,
		FirstName:   getUserByIDRow.FirstName,
		LastName:    getUserByIDRow.LastName,
		Email:       getUserByIDRow.Email,
		PhoneNumber: getUserByIDRow.PhoneNumber,
	}
	return user, nil
}

func (repo *UserRepository) CreateUser(user *domain.User) error {
	ctx := context.Background()

	_, err := repo.DB.User.Create().
		SetFirstName(user.FirstName).
		SetLastName(user.LastName).
		SetEmail(user.Email).
		SetPhoneNumber(user.PhoneNumber).
		Save(ctx)
	if err != nil {
		err = fmt.Errorf("[repository.UserRepository.CreateUser] failed: user = %+v, error = %w ", user, err)
		return err
	}

	return nil
}

func (repo *UserRepository) UpdateUser(user *domain.User, id int) error {
	ctx := context.Background()
	err := repo.DB.User.UpdateOneID(id).
		SetFirstName(user.FirstName).
		SetLastName(user.LastName).
		SetEmail(user.Email).
		SetPhoneNumber(user.PhoneNumber).
		Exec(ctx)
	if err != nil {
		err = fmt.Errorf("[repository.UserRepository.UpdateUser] failed: user = %+v, error = %w ", user, err)
		return err
	}
	return nil
}

func (repo *UserRepository) DeleteUserByID(id int) error {
	ctx := context.Background()
	err := repo.DB.User.DeleteOneID(id).Exec(ctx)
	if err != nil {
		err = fmt.Errorf("[repository.UserRepository.DeleteUserByID] failed: id = %d, error = %w ", id, err)
		return err
	}
	return nil
}
