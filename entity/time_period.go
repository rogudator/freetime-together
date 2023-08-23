package entity

type Period struct {
	UserID string `bson:"user_id"`
	Name string `bson:"name"`
	TimeFrom string `bson:"time_from"`
	TimeTo string `bson:"time_to"`
	Periodicity byte `bson:"periodicity"`
}