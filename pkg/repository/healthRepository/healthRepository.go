package healthRepository

import (
	"context"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"manntera.com/health-tracker-api/pkg/repository"
)

type HealthRepository struct {
	client         *firestore.Client
	userCollection *firestore.CollectionRef
}

var _ IHealthRepository = &HealthRepository{}

func NewHealthRepository(ctx context.Context) (*HealthRepository, error) {
	client, err := firestore.NewClient(ctx, repository.PROJECT_ID)

	if err != nil {
		return nil, err
	}

	col := client.Collection(repository.COLLECTION_USER)

	repo := HealthRepository{client: client, userCollection: col}

	return &repo, nil
}

func (repo *HealthRepository) GetHealthCollectionFromUserId(userId string) *firestore.CollectionRef {
	return repo.userCollection.Doc(userId).Collection(repository.COLLECTION_HEALTH)
}

func (repo *HealthRepository) AddData(ctx context.Context, userId string, timestamp int64, healthScore int, comment string, medicineName string) (*Health, error) {
	doc, _, addErr := repo.GetHealthCollectionFromUserId(userId).Add(ctx, map[string]interface{}{})

	if addErr != nil {
		return nil, addErr
	}
	healthData := Health{
		Id:           doc.ID,
		Timestamp:    timestamp,
		HealthScore:  healthScore,
		Comment:      comment,
		MedicineName: medicineName,
	}

	_, err := doc.Set(ctx, healthData)

	if err != nil {
		return nil, err
	}

	return &healthData, nil
}

func (repo *HealthRepository) GetDataFromTime(ctx context.Context, userId string, startTime int64, endTime int64) ([]Health, error) {
	healthCollection := repo.GetHealthCollectionFromUserId(userId)
	iter := healthCollection.Query.
		Where("Timestamp", ">=", startTime).
		Where("Timestamp", "<=", endTime).
		OrderBy("Timestamp", firestore.Asc).
		Documents(ctx)

	var Healths []Health
	for {
		doc, err := iter.Next()
		if err != nil {
			if err == iterator.Done {
				break
			}
			return nil, err
		}
		var health Health
		if err := doc.DataTo(&health); err != nil {
			return nil, err
		}
		health.Id = doc.Ref.ID
		Healths = append(Healths, health)
	}
	return Healths, nil
}

func (repo *HealthRepository) GetDataFromUuid(ctx context.Context, userId string, uuid string) (*Health, error) {
	doc, err := repo.GetHealthCollectionFromUserId(userId).Doc(uuid).Get(ctx)

	if err != nil {
		return nil, err
	}

	var health Health
	if err := doc.DataTo(&health); err != nil {
		return nil, err
	}

	return &health, nil
}

func (repo *HealthRepository) DeleteData(ctx context.Context, userId string, uuid string) error {
	_, err := repo.GetHealthCollectionFromUserId(userId).Doc(uuid).Delete(ctx)

	if err != nil {
		return err
	}

	return nil
}
