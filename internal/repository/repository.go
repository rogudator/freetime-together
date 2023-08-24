package repository

import "go.mongodb.org/mongo-driver/mongo"


type Repository struct {
	TimePeriodsList
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{TimePeriodsList: NewTimePeriodsListMongo(db)}
}