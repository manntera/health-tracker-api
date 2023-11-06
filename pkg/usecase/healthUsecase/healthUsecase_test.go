package healthUsecase

import (
	"context"
	"reflect"
	"testing"

	"go.uber.org/mock/gomock"
	"manntera.com/health-tracker-api/pkg/repository/healthRepository"
	mock_healthRepository "manntera.com/health-tracker-api/pkg/repository/healthRepository/mock"
	"manntera.com/health-tracker-api/pkg/repository/userRepository"
	mock_userRepository "manntera.com/health-tracker-api/pkg/repository/userRepository/mock"
)

func TestGetHealth(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ctx := context.Background()

	mock_healthRepo := mock_healthRepository.NewMockIHealthRepository(mockCtrl)
	health := healthRepository.Health{
		Id:          "e99f89b3",
		HealthScore: 2,
		Comment:     "comment1",
		Timestamp:   1766169806,
	}

	mock_healthRepo.EXPECT().GetDataFromTime(ctx, "user2", int64(1700000000), int64(1800000000)).Return([]healthRepository.Health{health}, nil)

	mock_userRepo := mock_userRepository.NewMockIUserRepository(mockCtrl)
	mock_userRepo.EXPECT().GetData(ctx, "user2").Return(&userRepository.User{Id: "user2"}, nil)

	HealthUsecase := NewHealthUsecase(mock_healthRepo, mock_userRepo)
	result, err := HealthUsecase.GetData(ctx, "user2", int64(1700000000), int64(1800000000))
	if err != nil {
		t.Errorf("TestGetData failed, expected error: %v, got: %v", nil, err)
	}
	if !reflect.DeepEqual(result, []healthRepository.Health{health}) {
		t.Errorf("TestGetData failed, expected result: %v, got: %v", []healthRepository.Health{health}, result)
	}
}
