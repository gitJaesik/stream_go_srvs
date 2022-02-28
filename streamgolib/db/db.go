package db

import (
	"context"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/gitJaesik/stream_go_srvs/streamgolib/logger"
	"golang.org/x/crypto/bcrypt"
)

type StreamGoLibDB interface {
	MongoPing(ctx context.Context) error
	ListSomething(ctx context.Context) (interface{}, error)
	RegisterAndCheckIndexes(ctx context.Context) error
}

func GenerateSomeKey(id, someKey string, date time.Time) string {
	encodeKey := fmt.Sprintf("%d/%02d/%02d/%s/%s", date.Year(), date.Month(), date.Day(), someKey, id)
	return base64.RawURLEncoding.EncodeToString([]byte(encodeKey))
}

func ParseSomeKey(keyKey string) (string, error) {
	decodeKey, err := base64.RawURLEncoding.DecodeString(keyKey)
	if err != nil {
		return "", err
	}
	return string(decodeKey), nil
}

func HashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logger.Logger.Infow("HashPassword", "hashed password err", err)
		return "", err
	}
	return string(hashPassword), nil
}
