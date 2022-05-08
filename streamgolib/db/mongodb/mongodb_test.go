package mongodb

import (
	"context"
	"testing"

	pbsas "github.com/gitJaesik/stream_go_srvs/streamgolib/gen/proto/go/stream_auth_server/v1"
	"github.com/gitJaesik/stream_go_srvs/streamgolib/logger"
)

func init() {
}

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
