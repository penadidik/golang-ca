package usecase_test

import (
	"context"
	"errors"
	"testing"
	"user-service/model"
	"github.com/penadidik/golang-ca/user-service/repository/mocks"
	"user-service/usecase"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUser(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	uc := usecase.NewUserUsecase(mockRepo)

	user := &model.User{Name: "Alice", Email: "alice@example.com"}

	mockRepo.On("CreateUser", mock.Anything, user).Return(user, nil)

	result, err := uc.CreateUser(context.TODO(), user)
	assert.NoError(t, err)
	assert.Equal(t, user, result)
	mockRepo.AssertExpectations(t)
}
