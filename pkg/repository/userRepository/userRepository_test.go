package userRepository

import (
	"context"
	"log"
	"testing"
)

const testId = "123"

func Initialize(t *testing.T) {
	t.Setenv("FIRESTORE_EMULATOR_HOST", "localhost:8080")
}
func TestAddData(t *testing.T) {
	Initialize(t)
	ctx := context.Background()
	repo, err := NewUserRepository(ctx)
	if err != nil {
		t.Errorf("Error creating user repository: %v", err)
	}

	t.Run("Test AddUser", func(t *testing.T) {
		userData := &User{
			Id:    testId,
			Email: "studio@manntera.com",
			Name:  "test",
		}
		err := repo.AddData(ctx, userData)
		if err != nil {
			t.Errorf("Error adding data: %v", err)
		}
	})
}

func TestGetData(t *testing.T) {
	Initialize(t)

	ctx := context.Background()
	repo, err := NewUserRepository(ctx)
	if err != nil {
		t.Errorf("Error creating user repository: %v", err)
	}

	t.Run("Test GetUser", func(t *testing.T) {
		user, err := repo.GetData(ctx, testId)
		if err != nil {
			t.Errorf("Error getting data: %v", err)
		}
		log.Print("【GetResult】 ", user)
		if user.Id != testId {
			t.Errorf("Expected id to be %s but got %v", testId, user.Id)
		}
		if user.Email != "studio@manntera.com" {
			t.Errorf("Expected email to be studio@manntera.com but got %v", user.Email)
		}
		if user.Name != "test" {
			t.Errorf("Expected name to be test but got %v", user.Name)
		}
	})
}

func TestDeleteData(t *testing.T) {
	Initialize(t)

	ctx := context.Background()
	repo, err := NewUserRepository(ctx)
	if err != nil {
		t.Errorf("Error creating user repository: %v", err)
	}

	t.Run("Test DeleteUser", func(t *testing.T) {
		err := repo.DeleteData(ctx, testId)
		if err != nil {
			t.Errorf("Error deleting data: %v", err)
		}
	})
}
