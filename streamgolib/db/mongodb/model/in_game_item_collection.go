package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InGameItemDocument struct {
	ID                  primitive.ObjectID `bson:"_id"`
	InGameItemId        int32              `bson:"inGameItemId"`
	Name                string             `bson:"name"`
	SpawnAbleStageLevel string             `bson:"spawnAbleStageLevel"`
	Description         string             `bson:"description"`
	CreatedAt           time.Time          `bson:"createdAt"`
	UpdatedAt           time.Time          `bson:"updatedAt"`
}
