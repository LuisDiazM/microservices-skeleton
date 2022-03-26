package database

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/LuisDiazM/firestore-reports/domain"
	"github.com/LuisDiazM/firestore-reports/domain/models"
	"google.golang.org/api/iterator"
)

const (
	COLLECTION_USER_ANSWERS = "users_report"
)

type DatabaseGatewayImp struct {
	client *firestore.Client
	ctx    context.Context
}

func NewDatabaseGatewayImp() domain.DatabaseGateway {
	return &DatabaseGatewayImp{}
}

func (firestoreImp *DatabaseGatewayImp) Setup() {
	var err error
	firestoreImp.client, err = firestore.NewClient(context.Background(), os.Getenv("GCP_PROJECT_ID"))
	if err != nil {
		log.Fatalln(err)
	}
	firestoreImp.ctx = context.TODO()
}

func (firestoreImp *DatabaseGatewayImp) GetUsersResponses() (*[]models.UserResponsesDTO, error) {
	var userResponses []models.UserResponsesDTO
	userResponsesRef := firestoreImp.client.Collection(COLLECTION_USER_ANSWERS)
	var data models.UserResponsesDTO
	iter := userResponsesRef.Documents(firestoreImp.ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		err = doc.DataTo(&data)
		if err != nil {
			return nil, err
		}
		userResponses = append(userResponses, data)
	}
	return &userResponses, nil
}
