package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DiceSkillDocument struct {
	ID           primitive.ObjectID `bson:"_id"`
	SkillId      int32              `bson:"skillId"`
	Name         string             `bson:"name"`
	SkillPower   int32              `bson:"skillPower"`
	SplashDamage float32            `bson:"splashDamage"`
	AttackSpeed  float32            `bson:"attackSpeed"`
	ReduceDamage float32            `bson:"reduceDamage"`
	PowerUp      float32            `bson:"powerUp"`
	AttackDesc   string             `bson:"attackDesc"`
	CreatedAt    time.Time          `bson:"createdAt"`
	UpdatedAt    time.Time          `bson:"updatedAt"`
}
