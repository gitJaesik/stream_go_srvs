package mongodb

import (
	"context"
	"errors"
	"strconv"
	"sync"

	"github.com/gitJaesik/grpcgolib/config"
	"github.com/gitJaesik/grpcgolib/db"
	"github.com/gitJaesik/grpcgolib/db/mongodb/model"
	"github.com/gitJaesik/grpcgolib/logger"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Set : for wire set
var Set = wire.NewSet(
	NewMongoDB,
)

// MongoDBClient ...
type MongoDBClient struct {
}

var once sync.Once
var mdbClient *MongoDBClient

func init() {
	ctx := context.Background()

	sglMongoDB := NewMongoDB()
	err := sglMongoDB.MongoPing(ctx)
	if err != nil {
		logger.Logger.Errorw("mongodb.go", "mongodb ping error", "server shutdown")
		panic(err)
	}
	err = sglMongoDB.RegisterAndCheckIndexes(ctx)
	if err != nil {
		logger.Logger.Errorw("mongodb.go", "mongodb index check error", "server shutdown")
		panic(err)
	}
}

// NewMongoDB ...
func NewMongoDB() db.StreamGoLibDB {
	once.Do(func() {
		logger.Logger.Info("NewMongoDB", "sync.Once.Do", "single ton")
		mdbClient = &MongoDBClient{}
	})
	if mdbClient == nil {
		panic(errors.New("mongodb Client error"))
	}
	return mdbClient
}

// MongoPing ...
func (mc *MongoDBClient) MongoPing(ctx context.Context) error {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.SglConfig.GetMongoUri()))
	if err != nil {
		logger.Logger.Errorw("MongoPing", "mongo connect error", err)
		return err
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		logger.Logger.Errorw("MongoPing", "mongo connect ping error", err)
		return err
	}
	logger.Logger.Infow("MongoPing", "mongo connect success", "Successfully connected and pinged.")
	return nil
}

// ListSomething ...
func (mc *MongoDBClient) ListSomething(ctx context.Context) (interface{}, error) {
	me := mongoExecutor{}

	result, err := me.execute(
		func(ctx context.Context, database *mongo.Database) (interface{}, error) {
			noteCollection := database.Collection(config.SglConfig.GetMongoNoteCollection())

			cursor, err := noteCollection.Find(ctx, bson.M{})
			if err != nil {
				logger.Logger.Errorw("ListSomething", "find note collection error", err)
				return nil, err
			}

			defer cursor.Close(ctx)

			// select
			notes := []model.NoteDocument{}
			for cursor.Next(ctx) {
				var note model.NoteDocument
				if err = cursor.Decode(&note); err != nil {
					logger.Logger.Errorw("ListSomething", "if err = cursor.Decode(&note); err != nil", err)
					return nil, err
				}
				notes = append(notes, note)
			}
			logger.Logger.Infow("ListSomething", "notes", notes)

			// response
			return notes, nil
		})
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (mc *MongoDBClient) RegisterAndCheckIndexes(ctx context.Context) error {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.SglConfig.GetMongoUri()))
	if err != nil {
		logger.Logger.Errorw("RegisterAndCheckIndexes", "mongo connect error", err)
		return err
	}
	defer client.Disconnect(ctx)

	// use database
	database := client.Database(config.SglConfig.GetMongoDatabase())
	userCollection := database.Collection(config.SglConfig.GetMongoUserCollection())

	indexes := []string{}
	indexes = append(indexes, "email")
	indexes = append(indexes, "username")

	for _, idxStr := range indexes {
		ret, err := EnsureIndexExist(userCollection, idxStr)
		if err != nil {
			return err
		}
		logger.Logger.Info("RegisterAndCheckIndexes", idxStr, ret)
		if !ret {
			err = CreateIndex(userCollection, true, false, -1, idxStr)
			if err != nil {
				logger.Logger.Errorw("RegisterAndCheckIndexes", "CreateIndex error", err)
				return err
			}
			logger.Logger.Info("RegisterAndCheckIndexes", "CreateIndex success", idxStr)
		}
	}

	ttlIndexes := []string{}
	ttlIndexes = append(ttlIndexes, "session.jwtsStatus")

	for _, idxStr := range ttlIndexes {
		ret, err := EnsureIndexExist(userCollection, idxStr)
		if err != nil {
			return err
		}
		logger.Logger.Info("RegisterAndCheckIndexes", "idxStr "+idxStr, ret)
		if !ret {
			jwtTTL, err := strconv.ParseInt(config.SglConfig.GetJwtTtl(), 10, 64)
			if err != nil {
				logger.Logger.Errorw("RegisterAndCheckIndexes", "strconv parse int jwt error", err)
				return err
			}

			err = CreateIndex(userCollection, false, true, int32(jwtTTL), idxStr)
			if err != nil {
				logger.Logger.Errorw("RegisterAndCheckIndexes", "CreateIndex error", err)
				return err
			}
			logger.Logger.Info("RegisterAndCheckIndexes", "CreateIndex success", idxStr)
		}
	}

	return nil
}
