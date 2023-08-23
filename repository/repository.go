package repository

import "go.mongodb.org/mongo-driver/mongo"


type repository struct {
	TimePeriodsList
}

func NewRepository(db *mongo.Database) *repository {
	return &repository{TimePeriodsList: NewTimePeriodsListMongo(db)}
}