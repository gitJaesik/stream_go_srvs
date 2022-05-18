package mongodb

import (
	"context"
	"errors"
	"testing"

	pbsas "github.com/gitJaesik/stream_go_srvs/streamgolib/gen/proto/go/stream_auth_server/v1"
	"github.com/gitJaesik/stream_go_srvs/streamgolib/logger"
)

func init() {
}

// var emotions []primitive.ObjectID
var emotions []int

// Add remote all mongodb

// TestDropAll ...
func TestDropAll(t *testing.T) {
	sglMongoDb := NewMongoDB()
	if err := sglMongoDb.DropAll(context.Background()); err != nil {
		t.Fatalf(`drop all error :%q`, err)
		panic(err)
	}
	t.Log("mongodb drop collections success")
}

/*
// TestPanicEmotion ...
func TestPanicTest(t *testing.T) {
	logger.Logger.Infow("TestPanicEmotion", "emotions", emotions)
	err := errors.New("test end")
	panic(err)
}
*/

// TestCreateIndex ...
func TestCreateIndex(t *testing.T) {

	sglMongoDb := NewMongoDB()

	err := sglMongoDb.RegisterAndCheckIndexes(context.Background())
	if err != nil {
		t.Fatalf(`mongodb register and check indexes error :%q`, err)
		panic(err)
	}
	t.Log("mongodb register and check indexes success")
}

// TestInsertSometing ...
func TestInsertSometing(t *testing.T) {
	ctx := context.Background()
	// ctxWithId := context.WithValue(ctx, stream_authentication.UserIDKey{}, gUserIDAdmin)
	sglMongoDb := NewMongoDB()

	// gpir := &pbsas.GetPlayerInfoRequest{}

	ret, err := sglMongoDb.InsertSomething(ctx)
	if err != nil {
		t.Fatalf(`InsertSomething error %v`, err)
	}

	logger.Logger.Infow("TestInsertSometing", "ret", ret)

}

// 사전 데이타 삽입

// Emotion

// TestCreateEmotionDocOne ...
func TestCreateEmotionDocOne(t *testing.T) {
	ctx := context.Background()
	// ctxWithId := context.WithValue(ctx, stream_authentication.UserIDKey{}, gUserIDAdmin)
	sglMongoDb := NewMongoDB()
	ret, err := sglMongoDb.CreateEmotion(ctx, &pbsas.Emotion{
		EmotionId:     1,
		Name:          "따봉",
		Description:   "쫌 하는데?",
		PurchaseGold:  100,
		PurchaseJewel: 300,
	})
	if err != nil {
		t.Fatalf(`TestCreateEmotionDocOne error %v`, err)
	}

	// objectid object id
	// emotions = append(emotions, ret.(*mongo.InsertOneResult).InsertedID.(primitive.ObjectID))
	logger.Logger.Infow("TestCreateEmotionDocOne ", "ret", ret)
}

// // TestPanicEmotion ...
// func TestPanicEmotion(t *testing.T) {
// 	logger.Logger.Infow("TestPanicEmotion", "emotions", emotions)
// 	err := errors.New("test end")
// 	panic(err)
// }

// TestCreateEmotionDocTwo ...
func TestCreateEmotionDocTwo(t *testing.T) {
	ctx := context.Background()
	sglMongoDb := NewMongoDB()
	ret, err := sglMongoDb.CreateEmotion(ctx, &pbsas.Emotion{
		EmotionId:     2,
		Name:          "비웃는",
		Description:   "너님 한참 멀었음",
		PurchaseGold:  100,
		PurchaseJewel: 200,
	})
	if err != nil {
		t.Fatalf(`TestCreateEmotionDocTwo error %v`, err)
	}
	logger.Logger.Infow("TestCreateEmotionDocTwo ", "ret", ret)
}

// TestCreateEmotionDocThree ...
func TestCreateEmotionDocThree(t *testing.T) {
	ctx := context.Background()
	sglMongoDb := NewMongoDB()
	ret, err := sglMongoDb.CreateEmotion(ctx, &pbsas.Emotion{
		EmotionId:     3,
		Name:          "언어",
		Description:   "그거밖에 안되냐?",
		PurchaseGold:  100,
		PurchaseJewel: 250,
	})
	if err != nil {
		t.Fatalf(`TestCreateEmotionDocThree error %v`, err)
	}
	logger.Logger.Infow("TestCreateEmotionDocThree ", "ret", ret)
}

// TestCreateEmotionDocFour ...
func TestCreateEmotionDocFour(t *testing.T) {
	ctx := context.Background()
	sglMongoDb := NewMongoDB()
	ret, err := sglMongoDb.CreateEmotion(ctx, &pbsas.Emotion{
		EmotionId:     4,
		Name:          "ㅈ밥",
		Description:   "아이언이냐?",
		PurchaseGold:  100,
		PurchaseJewel: 250,
	})
	if err != nil {
		t.Fatalf(`TestCreateEmotionDocFour error %v`, err)
	}
	logger.Logger.Infow("TestCreateEmotionDocFour ", "ret", ret)
}

// TestCreateEmotionDocFive ...
func TestCreateEmotionDocFive(t *testing.T) {
	ctx := context.Background()
	sglMongoDb := NewMongoDB()
	ret, err := sglMongoDb.CreateEmotion(ctx, &pbsas.Emotion{
		EmotionId:     5,
		Name:          "Respect",
		Description:   "잘하네",
		PurchaseGold:  100,
		PurchaseJewel: 100,
	})
	if err != nil {
		t.Fatalf(`TestCreateEmotionDocFive error %v`, err)
	}
	logger.Logger.Infow("TestCreateEmotionDocFive ", "ret", ret)
}

// TODO: make player for 송정헌, 장영인, 피재식

// TestCreatePlayerInfoEmptyDoc ...
func TestCreatePlayerInfoEmptyDoc(t *testing.T) {
	ctx := context.Background()
	// ctxWithId := context.WithValue(ctx, stream_authentication.UserIDKey{}, gUserIDAdmin)
	sglMongoDb := NewMongoDB()

	// gpir := &pbsas.GetPlayerInfoRequest{}

	ret, err := sglMongoDb.CreatePlayerInfo(ctx, &pbsas.CreatePlayerInfoRequest{
		NickName:       "SuperCell",
		Id:             1,
		PlayerLevel:    1,
		Tier:           "골드",
		WinningRate:    0,
		VictoryCount:   0,
		DefeatCount:    0,
		MaxWave:        3,
		TrophyCount:    30,
		OwningJewel:    300,
		OwningGold:     100,
		OwningEmoticon: []int32{1, 2, 3, 4, 5},
	})
	if err != nil {
		t.Fatalf(`CreatePlayerInfo error %v`, err)
	}

	logger.Logger.Infow("TestCreatePlayerInfo ", "ret", ret)

}

// TestCreatePlayerInfo ...
func TestCreatePlayerInfo(t *testing.T) {
	ctx := context.Background()
	// ctxWithId := context.WithValue(ctx, stream_authentication.UserIDKey{}, gUserIDAdmin)
	sglMongoDb := NewMongoDB()

	// gpir := &pbsas.GetPlayerInfoRequest{}

	ret, err := sglMongoDb.CreatePlayerInfo(ctx, &pbsas.CreatePlayerInfoRequest{})
	if err != nil {
		t.Fatalf(`CreatePlayerInfo error %v`, err)
	}

	logger.Logger.Infow("TestCreatePlayerInfo ", "ret", ret)

}

// func TestInitMongo(t *testing.T) {
// 	sglMongoDb := NewMongoDB()

// 	err := sglMongoDb.
// }

func TestPanic(t *testing.T) {
	err := errors.New("test end")
	panic(err)
}
