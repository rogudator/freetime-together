package entity

type Period struct {
	UserID      string `bson:"user_id" json:"user_id"`
	Name        string `bson:"name" json:"name"`
	TimeFrom    string `bson:"time_from" json:"time_from"`
	TimeTo      string `bson:"time_to" json:"time_to"`
	Periodicity byte   `bson:"periodicity" json:"periodicity"`
}