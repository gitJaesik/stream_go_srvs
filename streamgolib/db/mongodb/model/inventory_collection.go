package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InventoryDocument struct {
	ID                  primitive.ObjectID   `bson:"_id"`
	InventoryId         int32                `bson:"inventoryId"`
	TotalInventoryCount int32                `bson:"totalInventoryCount"`
	Boss                primitive.ObjectID   `bson:"boss"`
	PlayerSkillID       primitive.ObjectID   `bson:"playerSkillID"`
	DiceNum             int32                `bson:"diceNum"`
	DiceIDList          []primitive.ObjectID `bson:"diceIDList"`
	OwningEmoticons     []primitive.ObjectID `bson:"owningEmoticons"`
	CreatedAt           time.Time            `bson:"createdAt"`
	UpdatedAt           time.Time            `bson:"updatedAt"`
}
