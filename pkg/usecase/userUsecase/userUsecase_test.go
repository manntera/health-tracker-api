package userUsecase

import (
	"context"
	"reflect"
	"testing"

	"go.uber.org/mock/gomock"
	"manntera.com/health-tracker-api/pkg/repository/userRepository"
	mock_userRepository "manntera.com/health-tracker-api/pkg/repository/userRepository/mock"
)

func TestGetUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ctx := context.Background()

	mock := mock_userRepository.NewMockIUserRepository(mockCtrl)
	user := userRepository.User{
		Id:    "user1",
		Email: "studio@manntera.com",
		Name:  "user1",
	}
	mock.EXPECT().GetData(ctx, "user1").Return(&user, nil)

	UserUsecase := NewUserUsecase(mock)
	result, err := UserUsecase.GetUser(ctx, "user1")
	if err != nil {
		t.Errorf("TestGetData failed, expected error: %v, got: %v", nil, err)
	}
	if !reflect.DeepEqual(result, &user) {
		t.Errorf("TestGetData failed, expected result: %v, got: %v", &user, result)
	}
}
