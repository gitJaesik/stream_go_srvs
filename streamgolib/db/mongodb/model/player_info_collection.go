package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PlayerInfoDocument struct {
	ID             primitive.ObjectID `bson:"_id"`
	NickName       string             `bson:"nickName" validate:"required,min=2,max=20"`
	PlayerId       int32              `bson:"playerId"`
	PlayerLevel    int32              `bson:"playerLevel"`
	Tier           string             `bson:"tier"`
	WinningRate    float32            `bson:"winningRate"`
	VictoryCount   int32              `bson:"victoryCount"`
	DefeatCount    int32              `bson:"defeatCount"`
	MaxWave        int32              `bson:"maxWave"`
	TrophyCount    int32              `bson:"trophyCount"`
	OwningJewel    int32              `bson:"owningJewel"`
	OwningGold     int32              `bson:"owningGold"`
	OwningBox      int32              `bson:"owningBox"`
	OwningMap      int32              `bson:"owningMap"`
	OwningEmoticon []int32            `bson:"owningEmoticon"`
	OwningDice     []int32            `bson:"owningDice"`
	DailyGoldLevel int32              `bson:"dailyGoldLevel"`
	CreatedAt      time.Time          `bson:"createdAt"`
	UpdatedAt      time.Time          `bson:"updatedAt"`
}
