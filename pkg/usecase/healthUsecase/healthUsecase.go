package healthUsecase

import (
	"context"
	"errors"

	"manntera.com/health-tracker-api/pkg/repository/healthRepository"
	"manntera.com/health-tracker-api/pkg/repository/userRepository"
)

type HealthUsecase struct {
	healthRepo healthRepository.IHealthRepository
	userRepo   userRepository.IUserRepository
}

func NewHealthUsecase(healthRepo healthRepository.IHealthRepository, userRepo userRepository.IUserRepository) *HealthUsecase {
	return &HealthUsecase{healthRepo: healthRepo, userRepo: userRepo}
}

func (u *HealthUsecase) AddData(ctx context.Context, userId string, healthScore int, comment string, timestamp int64, medicineName string) (*healthRepository.Health, error) {
	user, userDataErr := u.userRepo.GetData(ctx, userId)
	if userDataErr != nil {
		return nil, userDataErr
	}

	health, err := u.healthRepo.AddData(ctx, user.Id, timestamp, healthScore, comment, medicineName)
	if err != nil {
		return nil, err
	}
	return health, nil
}

func (u *HealthUsecase) GetData(ctx context.Context, userId string, startTime, endTime int64) ([]healthRepository.Health, error) {
	user, userDataErr := u.userRepo.GetData(ctx, userId)
	if userDataErr != nil {
		return nil, userDataErr
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	healthData, err := u.healthRepo.GetDataFromTime(ctx, user.Id, startTime, endTime)
	if err != nil {
		return nil, err
	}
	return healthData, nil
}

func (u *HealthUsecase) DeleteData(ctx context.Context, userId string, uuid string) (*healthRepository.Health, error) {
	user, userDataErr := u.userRepo.GetData(ctx, userId)
	if userDataErr != nil {
		return nil, userDataErr
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	healthData, healthGetErr := u.healthRepo.GetDataFromUuid(ctx, user.Id, uuid)
	if healthGetErr != nil {
		return nil, healthGetErr
	}
	if healthData == nil {
		return nil, errors.New("health data not found")
	}

	err := u.healthRepo.DeleteData(ctx, user.Id, healthData.Id)
	if err != nil {
		return nil, err
	}
	return healthData, nil
}
