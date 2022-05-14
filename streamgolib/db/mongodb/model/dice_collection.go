package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DiceDocument struct {
	ID                 primitive.ObjectID   `bson:"_id"`
	DiceId             int32                `bson:"diceId"`
	Name               string               `bson:"name"`
	Description        string               `bson:"description"`
	DiceClassLevel     int32                `bson:"diceClassLevel"`
	BuyablePlayerLevel int32                `bson:"buyablePlayerLevel"`
	PurchaseGold       int32                `bson:"purchaseGold"`
	PurchaseJewel      int32                `bson:"purchaseJewel"`
	QuaterCount        int32                `bson:"quaterCount"`
	QuaterLevelMax     int32                `bson:"quaterLevelMax"`
	DiceLevel          int32                `bson:"diceLevel"`
	IsEquip            bool                 `bson:"isEquip"`
	IsNewDice          bool                 `bson:"isNewDice"`
	AttackType         primitive.ObjectID   `bson:"attackType"`
	AttackTarget       string               `bson:"attackTarget"`
	OwningSKillIDs     []primitive.ObjectID `bson:"owningSKillIDs"`
	CriticalDamage     int32                `bson:"criticalDamage"`
	CreatedAt          time.Time            `bson:"createdAt"`
	UpdatedAt          time.Time            `bson:"updatedAt"`
}
