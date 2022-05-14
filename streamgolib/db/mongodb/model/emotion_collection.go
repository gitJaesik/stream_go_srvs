package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmotionDocument struct {
	ID            primitive.ObjectID `bson:"_id"`
	Name          string             `bson:"name"`
	EmotionId     int32              `bson:"emotionId"`
	Description   string             `bson:"description"`
	PurchaseGold  int32              `bson:"purchaseGold"`
	PurchaseJewel int32              `bson:"purchaseJewel"`
	CreatedAt     time.Time          `bson:"createdAt"`
	UpdatedAt     time.Time          `bson:"updatedAt"`
}
