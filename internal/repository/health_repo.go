package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

const heartbeatCollection = "heartbeat"

type HealthRepo struct {
	dbCol *mongo.Collection
}

func NewHealthRepo(db *mongo.Database) *HealthRepo {
	return &HealthRepo{dbCol: db.Collection(heartbeatCollection)}
}

func (h *HealthRepo) InsertHeartbeat(timestamp time.Time) error {

	_, err := h.dbCol.InsertOne(context.Background(), bson.M{
		"timestamp": timestamp,
	})
	return err
}

func (h *HealthRepo) GetLatestHeartbeat() (time.Time, error) {
	var result struct {
		Timestamp time.Time `bson:"timestamp"`
	}

	err := h.dbCol.
		FindOne(context.Background(), bson.M{}).
		Decode(&result)
	if err != nil {
		return time.Time{}, err
	}

	return result.Timestamp, nil
}
