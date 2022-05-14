package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ShopMainDocument struct {
	ID                           primitive.ObjectID `bson:"_id"`
	DiceId                       int32              `bson:"diceId"`
	Name                         string             `bson:"name"`
	DailyGoldDisplayCount        int32              `bson:"dailyGoldDisplayCount"`
	DailyDiceProductDisplayCount int32              `bson:"dailyDiceProductDisplayCount"`
	EmoticonDisplayCount         int32              `bson:"emoticonDisplayCount"`
	RoyalBoxDispalyCount         int32              `bson:"royalBoxDispalyCount"`
	JewlyDispalyCount            int32              `bson:"jewlyDispalyCount"`
	GoldDispalyCount             int32              `bson:"goldDispalyCount"`
	CreatedAt                    time.Time          `bson:"createdAt"`
	UpdatedAt                    time.Time          `bson:"updatedAt"`
}
