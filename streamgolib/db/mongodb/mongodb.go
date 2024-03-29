package mongodb

import (
	"context"
	"errors"
	"strconv"
	"sync"
	"time"

	"github.com/gitJaesik/stream_go_srvs/streamgolib/config"
	"github.com/gitJaesik/stream_go_srvs/streamgolib/db"
	"github.com/gitJaesik/stream_go_srvs/streamgolib/db/mongodb/model"
	pbsas "github.com/gitJaesik/stream_go_srvs/streamgolib/gen/proto/go/stream_auth_server/v1"
	"github.com/gitJaesik/stream_go_srvs/streamgolib/logger"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// DropAll ...
func (mc *MongoDBClient) DropAll(ctx context.Context) error {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.SglConfig.GetMongoUri()))
	if err != nil {
		logger.Logger.Errorw("DropAll", "mongo connect error", err)
		return err
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	logger.Logger.Infow("DropAll", "client", client)
	dbName := config.SglConfig.GetMongoDatabase()
	logger.Logger.Infow("DropAll", "dbName", dbName)
	database := client.Database(dbName)
	logger.Logger.Infow("DropAll", "database", database)
	collections, err := database.ListCollectionNames(ctx, bson.M{}, nil)
	logger.Logger.Infow("DropAll", "collections", collections)
	for _, collection := range collections {
		logger.Logger.Infow("DropAll", "current collection", collection)
		if err := database.Collection(collection).Drop(ctx); err != nil {
			logger.Logger.Errorw("DropAll", collection, err)
			return err
		}
	}
	// if err := database.Drop(ctx); err != nil {
	// 	logger.Logger.Errorw("DropAll", "mongo drop error", err)
	// 	return err
	// }
	logger.Logger.Infow("DropAll", "drop collections success", "Successfully drops all collection")
	return nil
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

// InsertSomething ...
// func (mc *MongoDBClient) InsertSomething(ctx context.Context, req *pbsas.GetPlayerInfoRequest) (interface{}, error) {
func (mc *MongoDBClient) InsertSomething(ctx context.Context) (interface{}, error) {
	// logger.Logger.Infof("insertingReq: %v, ", req)

	// userID := ctx.Value(stream_authentication.UserIDKey{}).(string)
	// logger.Logger.Infow("InsertSomething", "", userID)
	// userSummery, err := mc.GetUserSummaryByUserIDKey(ctx, userID)
	// if err != nil {
	// 	panic(err)
	// }

	me := mongoExecutor{}

	result, err := me.execute(
		func(ctx context.Context, database *mongo.Database) (interface{}, error) {
			noteCollection := database.Collection(config.SglConfig.GetMongoNoteCollection())

			now := time.Now()

			note := model.NoteDocument{
				ID:        primitive.NewObjectID(),
				NoteKey:   "notekey",
				CreatedAt: now,
				UpdatedAt: now,
				// Publisher: *userSummary,
				Title:    "hello",
				Contents: model.ContentDocument{},
			}

			insertOption := options.InsertOne()

			result, err := noteCollection.InsertOne(ctx, note, insertOption)
			if err != nil {
				return nil, err
			}

			note.ID = result.InsertedID.(primitive.ObjectID)
			noteKey := db.GenerateSomeKey(note.ID.Hex(), note.Title, now)
			// filter update
			return noteKey, nil
		})
	if err != nil {
		return nil, err
	}

	return result, nil
}

// CreateEmotion ...
func (mc *MongoDBClient) CreateEmotion(ctx context.Context, packet *pbsas.Emotion) (interface{}, error) {
	me := mongoExecutor{}

	result, err := me.execute(
		func(ctx context.Context, database *mongo.Database) (interface{}, error) {
			emotionCollection := database.Collection(config.SglConfig.GetMongoEmotionCollection())

			now := time.Now()

			emotion := model.EmotionDocument{
				ID:            primitive.NewObjectID(),
				Name:          packet.Name,
				EmotionId:     packet.EmotionId,
				Description:   packet.Description,
				PurchaseGold:  packet.PurchaseGold,
				PurchaseJewel: packet.PurchaseJewel,
				CreatedAt:     now,
				UpdatedAt:     now,
			}
			insertOption := options.InsertOne()

			result, err := emotionCollection.InsertOne(ctx, emotion, insertOption)
			if err != nil {
				return nil, err
			}

			// filter update
			return result, nil
		})
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreatePlayerInfo ...
func (mc *MongoDBClient) CreatePlayerInfo(ctx context.Context, packet *pbsas.CreatePlayerInfoRequest) (interface{}, error) {
	me := mongoExecutor{}

	result, err := me.execute(
		func(ctx context.Context, database *mongo.Database) (interface{}, error) {
			playerCollection := database.Collection(config.SglConfig.GetMongoPlayerCollection())

			now := time.Now()

			// https://www.mongodb.com/basics/mongodb-auto-increment#:~:text=Although%20MongoDB%20does%20not%20support,the%20current%20unique%20identifier%20value.
			player := model.PlayerInfoDocument{
				ID:             primitive.NewObjectID(),
				PlayerId:       packet.Id,
				NickName:       packet.NickName,
				PlayerLevel:    packet.PlayerLevel,
				Tier:           packet.Tier,
				WinningRate:    packet.WinningRate,
				VictoryCount:   packet.VictoryCount,
				DefeatCount:    packet.DefeatCount,
				MaxWave:        packet.MaxWave,
				TrophyCount:    packet.TrophyCount,
				OwningJewel:    packet.OwningJewel,
				OwningGold:     packet.OwningGold,
				OwningBox:      packet.OwningBox,
				OwningMap:      packet.OwningMap,
				OwningEmoticon: packet.OwningEmoticon,
				DailyGoldLevel: packet.DailyGoldLevel,
			}
			insertOption := options.InsertOne()

			result, err := playerCollection.InsertOne(ctx, player, insertOption)
			if err != nil {
				return nil, err
			}

			player.ID = result.InsertedID.(primitive.ObjectID)
			playerKey := db.GenerateSomeKey(player.ID.Hex(), player.NickName, now)
			// filter update
			return playerKey, nil
		})
	if err != nil {
		return nil, err
	}

	return result, nil

}

// https://www.mongodb.com/blog/post/quick-start-golang--mongodb--data-aggregation-pipeline
// https://www.mongodb.com/blog/post/quick-start-golang--mongodb--modeling-documents-with-go-data-structures
// https://www.mongodb.com/blog/post/mongodb-go-driver-updated-to-version-104
// https://www.mongodb.com/docs/drivers/go/v1.8/fundamentals/crud/query-document/

// GetPlayerInfo ...
func (mc *MongoDBClient) GetPlayerInfo(ctx context.Context, packet *pbsas.GetPlayerInfoRequest) (interface{}, error) {
	me := mongoExecutor{}

	result, err := me.execute(
		func(ctx context.Context, database *mongo.Database) (interface{}, error) {
			playerCollection := database.Collection(config.SglConfig.GetMongoPlayerCollection())

			/*
				if err := connection.collection.FindOne(ctx, bson.M{"name": recipeName}).Decode(&recipe); err != nil {
					return alexa.Response{}, err
				}
				var results bson.M
				err := fsFiles.FindOne(ctx, bson.M{}).Decode(&results)
				if err != nil {
					log.Fatal(err)
				}
				// you can print out the results
				fmt.Println(results)
			*/
			// cursor, err := playerCollection.Find(ctx, bson.M{})
			// ret := playerCollection.FindOne(ctx, bson.M{"playerId", 1})
			// var results bson.M
			var playerInfoDoc model.PlayerInfoDocument
			// matchStage := bson.D{{"$match", bson.D{{"podcast", id}}}}
			// groupStage := bson.D{{"$group", bson.D{{"_id", "$podcast"}, {"total", bson.D{{"$sum", "$duration"}}}}}}

			// err := playerCollection.FindOne(ctx, bson.M{"playerId": packet.PlayerId}).Decode(&results)
			// err := playerCollection.FindOne(ctx, bson.D{{"playerId", bson.D{{"$eq", packet.PlayerId}}}}).Decode(&results)

			err := playerCollection.FindOne(ctx, bson.D{{"playerId", packet.PlayerId}}).Decode(&playerInfoDoc)
			// err := playerCollection.FindOne(ctx, bson.D{{"playerId": packet.PlayerId}}).Decode(&results)
			if err != nil {
				logger.Logger.Errorw("GetPlayerInfo", "error", err)
				return nil, err
			}

			logger.Logger.Infow("GetPlayerInfo", "playerInfoDoc", playerInfoDoc)

			// packet *pbsas.GetPlayerInfoRequest

			response := &pbsas.GetPlayerInfoResponse{
				NickName:       playerInfoDoc.NickName,
				PlayerId:       playerInfoDoc.PlayerId,
				PlayerLevel:    playerInfoDoc.PlayerLevel,
				Tier:           playerInfoDoc.Tier,
				WinningRate:    playerInfoDoc.WinningRate,
				VictoryCount:   playerInfoDoc.VictoryCount,
				DefeatCount:    playerInfoDoc.DefeatCount,
				MaxWave:        playerInfoDoc.MaxWave,
				TrophyCount:    playerInfoDoc.TrophyCount,
				OwningJewel:    playerInfoDoc.OwningJewel,
				OwningGold:     playerInfoDoc.OwningGold,
				OwningBox:      playerInfoDoc.OwningBox,
				OwningMap:      playerInfoDoc.OwningMap,
				OwningEmoticon: playerInfoDoc.OwningEmoticon,
				// OwningDice:     playerInfoDoc.OwningDice,
				DailyGoldLevel: playerInfoDoc.DailyGoldLevel,
				// CreatedAt:      playerInfoDoc.CreatedAt,
				// UpdatedAt:      playerInfoDoc.UpdatedAt,
			}

			// you can print out the results
			// fmt.Println(results)
			// if err != nil {
			// 	logger.Logger.Errorw("ListSomething", "find note collection error", err)
			// 	return nil, err
			// }

			// defer cursor.Close(ctx)

			// select
			// notes := []model.NoteDocument{}
			// for cursor.Next(ctx) {
			// 	var note model.NoteDocument
			// 	if err = cursor.Decode(&note); err != nil {
			// 		logger.Logger.Errorw("ListSomething", "if err = cursor.Decode(&note); err != nil", err)
			// 		return nil, err
			// 	}
			// 	notes = append(notes, note)
			// }
			// logger.Logger.Infow("ListSomething", "notes", notes)

			// response
			return response, nil
		})
	if err != nil {
		return nil, err
	}

	return result, nil
}
