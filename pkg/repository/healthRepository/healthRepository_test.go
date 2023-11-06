package healthRepository

import (
	"context"
	"testing"
)

const userName = "user2"

var testData = Health{
	HealthScore: -5,
	Comment:     "comment3",
	Timestamp:   1758304171,
}

func TestAddData(t *testing.T) {
	t.Setenv("FIRESTORE_EMULATOR_HOST", "localhost:8080")
	ctx := context.Background()
	healthRepository, err := NewHealthRepository(ctx)
	if err != nil {
		t.Error(err)
	}

	t.Run("Test AddData", func(t *testing.T) {
		_, err = healthRepository.AddData(ctx, userName, testData.Timestamp, testData.HealthScore, testData.Comment)
		if err != nil {
			t.Error(err)
		}
	})
}

func TestGetData(t *testing.T) {
	t.Setenv("FIRESTORE_EMULATOR_HOST", "localhost:8080")

	ctx := context.Background()
	healthRepository, err := NewHealthRepository(ctx)
	if err != nil {
		t.Error(err)
	}

	t.Run("Test GetData", func(t *testing.T) {
		result, err := healthRepository.GetDataFromTime(ctx, userName, 0, 2000000000)
		if err != nil {
			t.Error(err)
		}
		if result[0].Comment != testData.Comment {
			t.Errorf("result is not equal to testData\ngot: %v\nwant: %v", result, testData)
		}
		if result[0].HealthScore != testData.HealthScore {
			t.Errorf("result is not equal to testData\ngot: %v\nwant: %v", result, testData)
		}
		if result[0].Timestamp != testData.Timestamp {
			t.Errorf("result is not equal to testData\ngot: %v\nwant: %v", result, testData)
		}
	})
}

func TestDeleteData(t *testing.T) {
	t.Setenv("FIRESTORE_EMULATOR_HOST", "localhost:8080")

	ctx := context.Background()
	healthRepository, err := NewHealthRepository(ctx)
	if err != nil {
		t.Error(err)
	}

	t.Run("Test DeleteData", func(t *testing.T) {
		err := healthRepository.DeleteData(ctx, userName, "bjKHOBApFz62t4QXakgF")
		if err != nil {
			t.Error(err)
		}
	})
}
