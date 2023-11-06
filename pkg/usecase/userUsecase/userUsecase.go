package userUsecase

import (
	"context"
	"errors"

	"manntera.com/health-tracker-api/pkg/repository/userRepository"
)

type UserUsecase struct {
	repository userRepository.IUserRepository
}

func NewUserUsecase(repository userRepository.IUserRepository) *UserUsecase {
	return &UserUsecase{repository: repository}
}

func (u *UserUsecase) GetUser(ctx context.Context, userId string) (*userRepository.User, error) {
	userData, err := u.repository.GetData(ctx, userId)
	if err != nil {
		return nil, err
	}
	return userData, nil
}

func (u *UserUsecase) AddUser(ctx context.Context, userData *userRepository.User) error {
	err := u.repository.AddData(ctx, userData)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserUsecase) DeleteUser(ctx context.Context, userId string) error {
	user, getErr := u.repository.GetData(ctx, userId)
	if getErr != nil {
		return getErr
	}

	if user == nil {
		return errors.New("user not found")
	}

	deleteErr := u.repository.DeleteData(ctx, userId)
	if deleteErr != nil {
		return deleteErr
	}
	return nil
}
