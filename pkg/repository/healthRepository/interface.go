package healthRepository

import "context"

type IHealthRepository interface {
	AddData(ctx context.Context, userId string, timestamp int64, healthScore int, comment string, medicineName string) (*Health, error)
	GetDataFromTime(ctx context.Context, userId string, startTime int64, endTime int64) ([]Health, error)
	GetDataFromUuid(ctx context.Context, userId string, uuid string) (*Health, error)
	DeleteData(ctx context.Context, userId string, uuid string) error
}
