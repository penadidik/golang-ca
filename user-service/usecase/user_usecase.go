package usecase

import (
	"context"
	"user-service/model"
	"user-service/repository"
)

type UserUsecase interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetAllUsers(ctx context.Context) ([]model.User, error)
	GetUserByID(ctx context.Context, id string) (*model.User, error)
	UpdateUser(ctx context.Context, id string, user *model.User) error
	DeleteUser(ctx context.Context, id string) error
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{repo: repo}
}

func (u *userUsecase) CreateUser(ctx context.Context, user *model.User) error {
	return u.repo.Create(ctx, user)
}

func (u *userUsecase) GetAllUsers(ctx context.Context) ([]model.User, error) {
	return u.repo.FindAll(ctx)
}

func (u *userUsecase) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	return u.repo.FindByID(ctx, id)
}

func (u *userUsecase) UpdateUser(ctx context.Context, id string, user *model.User) error {
	return u.repo.Update(ctx, id, user)
}

func (u *userUsecase) DeleteUser(ctx context.Context, id string) error {
	return u.repo.Delete(ctx, id)
}
