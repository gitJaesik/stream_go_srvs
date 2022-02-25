package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// type IconDocument struct {
// 	ID  primitive.ObjectID `bson:"_id"`
// 	Url string             `bson:"url"`
// }

type UserDocument struct {
	ID        primitive.ObjectID   `bson:"_id"`
	Email     string               `bson:"email,omitempty" validate:"required,email,min=5,max=20"`
	Password  string               `bson:"password,omitempty" validate:"required,min=8,max=20"`
	Username  string               `bson:"username,omitempty" validate:"required,min=5,max=15"`
	CreatedAt time.Time            `bson:"createdAt"`
	UpdatedAt time.Time            `bson:"updatedAt"`
	Notes     []primitive.ObjectID `bson:"notes"`
	Session   UserSessionDocument  `bson:"session"`
}

type JwtStatusDocument struct {
	Jwt    string `bson:"jwt"`
	Status string `bson:"status"`
}

type UserSessionDocument struct {
	ExpiredAt time.Time           `bson:"expiredAt"`
	JwtStatus []JwtStatusDocument `bson:"jwtsStatus"`
}

type UserSummaryDocument struct {
	UserID   primitive.ObjectID `bson:"_id"`
	Username string             `bson:"username"`
}
