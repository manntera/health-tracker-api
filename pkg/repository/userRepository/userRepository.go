package userRepository

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"manntera.com/health-tracker-api/pkg/repository"
)

type UserRepository struct {
	client     *firestore.Client
	collection *firestore.CollectionRef
}

var _ IUserRepository = &UserRepository{}

func NewUserRepository(ctx context.Context) (*UserRepository, error) {
	client, err := firestore.NewClient(ctx, repository.PROJECT_ID)

	if err != nil {
		return nil, err
	}

	col := client.Collection(repository.COLLECTION_USER)

	repo := UserRepository{client: client, collection: col}

	return &repo, nil
}

func (repo *UserRepository) GetUserDocumentFromId(id string) *firestore.DocumentRef {
	return repo.collection.Doc(id)
}

func (repo *UserRepository) AddData(ctx context.Context, userData *User) error {
	doc := repo.GetUserDocumentFromId(userData.Id)
	result, err := doc.Set(ctx, userData)
	if err != nil {
		return err
	}
	log.Print(result)

	return nil
}

func (repo *UserRepository) GetData(ctx context.Context, id string) (*User, error) {
	doc := repo.GetUserDocumentFromId(id)
	result, err := doc.Get(ctx)
	if err != nil {
		return nil, err
	}
	var user User
	if err := result.DataTo(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) DeleteData(ctx context.Context, id string) error {
	doc := repo.GetUserDocumentFromId(id)
	_, deleteErr := doc.Delete(ctx)
	if deleteErr != nil {
		return deleteErr
	}
	return nil
}

func (repo *UserRepository) UpdateData(ctx context.Context, id string, userData *User) error {
	doc := repo.GetUserDocumentFromId(id)
	_, err := doc.Set(ctx, userData)
	if err != nil {
		return err
	}
	return nil
}
