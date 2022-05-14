package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AttackTypeDocument struct {
	ID           primitive.ObjectID `bson:"_id"`
	Name         string             `bson:"name"`
	AttackTypeId int32              `bson:"attackTypeId"`
	DicePower    int32              `bson:"dicePower"`
	SplashDamage float32            `bson:"splashDamage"`
	AttackSpeed  int32              `bson:"attackSpeed"`
	AttackDesc   string             `bson:"attackDesc"`
	CreatedAt    time.Time          `bson:"createdAt"`
	UpdatedAt    time.Time          `bson:"updatedAt"`
}
