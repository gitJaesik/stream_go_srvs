package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmotionDeckDocument struct {
	ID            primitive.ObjectID `bson:"_id"`
	EmotionDeckId int32              `bson:"emotionDeckId"`
	Name          string             `bson:"name"`
	EmotionId0    primitive.ObjectID `bson:"emotionId0"`
	EmotionId1    primitive.ObjectID `bson:"emotionId1"`
	EmotionId2    primitive.ObjectID `bson:"emotionId2"`
	EmotionId3    primitive.ObjectID `bson:"emotionId3"`
	EmotionId4    primitive.ObjectID `bson:"emotionId4"`
	EmotionId5    primitive.ObjectID `bson:"emotionId5"`
	CreatedAt     time.Time          `bson:"createdAt"`
	UpdatedAt     time.Time          `bson:"updatedAt"`
}
