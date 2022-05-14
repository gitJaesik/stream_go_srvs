package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DiceDeckDocument struct {
	ID         primitive.ObjectID `bson:"_id"`
	DiceDeckId int32              `bson:"diceDeckId"`
	Name       string             `bson:"name"`
	Dice0      primitive.ObjectID `bson:"dice0"`
	Dice1      primitive.ObjectID `bson:"dice1"`
	Dice2      primitive.ObjectID `bson:"dice2"`
	Dice3      primitive.ObjectID `bson:"dice3"`
	Dice4      primitive.ObjectID `bson:"dice4"`
	Dice5      primitive.ObjectID `bson:"dice5"`
	CreatedAt  time.Time          `bson:"createdAt"`
	UpdatedAt  time.Time          `bson:"updatedAt"`
}
