package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TODO: refactor purchase colums
// FIXME: seperate data and bill
type BossDocument struct {
	ID             primitive.ObjectID `bson:"_id"`
	BossId         int32              `bson:"bossId"`
	Name           string             `bson:"name"`
	Description    string             `bson:"description"`
	ClassLevel     int32              `bson:"classLevel"`
	PurchaseGold   int32              `bson:"purchaseGold"`
	PurchaseJewel  int32              `bson:"purchaseJewel"`
	QuaterCount    int32              `bson:"quaterCount"`
	QuaterLevelMax int32              `bson:"quaterLevelMax"`
	Level          int32              `bson:"level"`
	CreatedAt      time.Time          `bson:"createdAt"`
	UpdatedAt      time.Time          `bson:"updatedAt"`
}
