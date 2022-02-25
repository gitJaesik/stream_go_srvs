package mongodb

import (
	"context"
	"time"

	"github.com/gitJaesik/grpcgolib/config"
	"github.com/gitJaesik/grpcgolib/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

type mongoExecutor struct {
}

type executeFunc func(ctx context.Context, database *mongo.Database) (interface{}, error)

func (me *mongoExecutor) execute(ef executeFunc) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := config.GetMongoUri()
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		logger.Logger.Errorw("execute", "mongo connect error", err)
		return nil, err
	}

	defer client.Disconnect(ctx)

	dbName := config.GetMongoDatabase()
	database := client.Database(dbName)

	return ef(ctx, database)
}

type executeTxnFunc func(
	sessCtx context.Context,
	database *mongo.Database,
	wcMajorityCollectionOpts *options.CollectionOptions,
) (interface{}, error)

func (me *mongoExecutor) executeWithTransaction(etf executeTxnFunc) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := config.GetMongoUri()
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		logger.Logger.Errorw("executeWithTransaction",
			"mongo connect error", err)
		return nil, err
	}

	defer client.Disconnect(ctx)

	dbName := config.GetMongoDatabase()
	database := client.Database(dbName)

	wcMajority := writeconcern.New(
		writeconcern.WMajority(),
		writeconcern.WTimeout(10*time.Second),
	)
	wcMajorityCollectionOpts := options.Collection().SetWriteConcern(wcMajority)

	session, err := client.StartSession()
	if err != nil {
		return nil, err
	}
	defer session.EndSession(ctx)

	// Step 1: Define the callback that specifies the sequence
	// of perform inside the transaction

	// Step 2:
	// Start a session and run the callback using WithTransaction.

	result, err := session.WithTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {
		return etf(sessCtx, database, wcMajorityCollectionOpts)
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}
