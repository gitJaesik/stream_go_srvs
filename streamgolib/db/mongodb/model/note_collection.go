package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NoteDocument struct {
	ID        primitive.ObjectID  `bson:"_id"`
	NoteKey   string              `bson:"noteKey"`
	CreatedAt time.Time           `bson:"createdAt"`
	UpdatedAt time.Time           `bson:"updatedAt"`
	Publisher UserSummaryDocument `bson:"publisher" validate:"required"`
	Title     string              `bson:"title"`
	Contents  ContentDocument     `bson:"content"`
}

type ContentDocument struct {
	content []string `bson:"content"`
}
