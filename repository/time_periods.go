package repository

import (
	"context"

	"github.com/rogudator/freetime-together/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const timePeriods = "time_periods"

type TimePeriodsListMongo struct {
	db *mongo.Database
}

func NewTimePeriodsListMongo(db *mongo.Database) *TimePeriodsListMongo {
	return &TimePeriodsListMongo{db: db}
}

type TimePeriodsList interface {
	CreateTimePeriod(ctx context.Context, p entity.Period) error
	GetTimePeriod(ctx context.Context, userID, timeFrom, timeTo string) (entity.Period, error)
	GetAllTimePeriods(ctx context.Context, userID string) ([]entity.Period, error)
	UpdateTimePeriod(ctx context.Context, old, new entity.Period) error
	DeleteTimePeriod(ctx context.Context, p entity.Period) error
}

func (r *TimePeriodsListMongo) CreateTimePeriod(ctx context.Context, p entity.Period) error {
	_, err := r.db.Collection(timePeriods).InsertOne(ctx, p)
	if err != nil {
		return err
	}
	return nil
}

func (r *TimePeriodsListMongo) GetTimePeriod(ctx context.Context, userID, timeFrom, timeTo string) (entity.Period, error) {
	filter := bson.D{
		{Key: "user_id", Value: userID},
		{Key: "time_from", Value: timeFrom},
		{Key: "time_to", Value: timeTo},
	}
	var result entity.Period
	err := r.db.Collection(timePeriods).FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return entity.Period{}, err
	}
	return result, nil
}

func (r *TimePeriodsListMongo) GetAllTimePeriods(ctx context.Context, userID string) ([]entity.Period, error) {
	filter := bson.D{{Key: "user_id", Value: userID}}
	results := make([]entity.Period, 0)
	cursor, err := r.db.Collection(timePeriods).Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var elem entity.Period
		err = cursor.Decode(&elem)
		if err != nil {
			return nil, err
		}
		results = append(results, elem)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (r *TimePeriodsListMongo) UpdateTimePeriod(ctx context.Context, old, new entity.Period) error {
	filter := old
	update := bson.M{
		"$set": new,
	}
	_, err := r.db.Collection(timePeriods).UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (r *TimePeriodsListMongo) DeleteTimePeriod(ctx context.Context, p entity.Period) error {
	_, err := r.db.Collection(timePeriods).DeleteOne(ctx, p)
	if err != nil {
		return err
	}
	return nil
}
