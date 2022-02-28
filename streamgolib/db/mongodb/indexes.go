package mongodb

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/gitJaesik/stream_go_srvs/streamgolib/config"
	"github.com/gitJaesik/stream_go_srvs/streamgolib/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

// CreateUniqueIndex create unique index
func CreateUniqueIndex(collection *mongo.Collection, keys ...string) error {
	keysDoc := bsonx.Doc{}

	for _, key := range keys {
		if strings.HasPrefix(key, "-") {
			keysDoc = keysDoc.Append(strings.TrimLeft(key, "-"), bsonx.Int32(-1))
		} else {
			keysDoc = keysDoc.Append(key, bsonx.Int32(1))
		}
	}
	idxRet, err := collection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    keysDoc,
			Options: options.Index().SetUnique(true),
		},
		options.CreateIndexes().SetMaxTime(10*time.Second),
	)

	if err != nil {
		return err
	}

	logger.Logger.Infow("CreateUniqueIndex", "collection.Indexes().CreateOne", idxRet)

	return nil

}

// ListIndexes ...
func ListIndexes(collection *mongo.Collection) {
	duration := 10 * time.Second
	batchSize := int32(10)
	cur, err := collection.Indexes().List(
		context.Background(), &options.ListIndexesOptions{&batchSize, &duration})
	if err != nil {
		panic(err)
	}
	for cur.Next(context.Background()) {
		index := bson.D{}
		cur.Decode(&index)
		logger.Logger.Infow("ListIndexes", "index found", index)
	}

}

// EnsureIndexExist will ensure the index model provided is on the given collection.
func EnsureIndexExist(c *mongo.Collection, key string) (bool, error) {
	idxs := c.Indexes()

	cur, err := idxs.List(context.Background())
	if err != nil {
		return false, err
	}

	found := false
	for cur.Next(context.Background()) {
		d := make(map[string]interface{})

		if err := cur.Decode(d); err != nil {
			return false, err
		}
		m := d["key"].(map[string]interface{})
		if _, ok := m[key]; ok {
			found = true
			break
		}
	}

	if found {
		return true, nil
	}

	return false, nil
}

// CreateTTLIndex create ttl index
func CreateTTLIndex(collection *mongo.Collection, keys ...string) error {
	keysDoc := bsonx.Doc{}

	for _, key := range keys {
		if strings.HasPrefix(key, "-") {
			keysDoc = keysDoc.Append(strings.TrimLeft(key, "-"), bsonx.Int32(-1))
		} else {
			keysDoc = keysDoc.Append(key, bsonx.Int32(1))
		}
	}

	// for jwt collection ttl
	ttlTime, err := strconv.ParseInt(config.SglConfig.GetJwtTtl(), 10, 64)
	if err != nil {
		logger.Logger.Infow("CreateTTLIndex",
			"ttlTime, err := strconv.ParseInt(config.SglConfig.GetJwtTtl(), 10, 64) error", err)

		// default
		ttlTime = 10800
	}

	idxRet, err := collection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    keysDoc,
			Options: options.Index().SetExpireAfterSeconds(int32(ttlTime)),
		},
		options.CreateIndexes().SetMaxTime(10*time.Second),
	)

	if err != nil {
		return err
	}

	logger.Logger.Infow("CreateTTLIndex", "collection.Indexes().CreateOne", idxRet)

	return nil

}

// CreateIndex create index, unique or ttl
func CreateIndex(collection *mongo.Collection, useUnique bool, useTtl bool, ttl int32, keys ...string) error {

	// validation check
	if useTtl {
		if ttl <= 0 {
			return errors.New("ttl is not validated")
		}
	}

	keysDoc := bsonx.Doc{}

	for _, key := range keys {
		if strings.HasPrefix(key, "-") {
			keysDoc = keysDoc.Append(strings.TrimLeft(key, "-"), bsonx.Int32(-1))
		} else {
			keysDoc = keysDoc.Append(key, bsonx.Int32(1))
		}
	}

	var indexOption *options.IndexOptions

	if useUnique && useTtl {
		indexOption = options.Index().SetExpireAfterSeconds(ttl).SetUnique(true)
	} else if useUnique {
		indexOption = options.Index().SetUnique(true)
	} else if useTtl {
		indexOption = options.Index().SetExpireAfterSeconds(ttl)
	} else {
		return errors.New("unexpected function callling")
	}

	idxRet, err := collection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    keysDoc,
			Options: indexOption,
		},
		options.CreateIndexes().SetMaxTime(10*time.Second),
	)

	if err != nil {
		return err
	}

	logger.Logger.Infow("CreateTTLIndex", "collection.Indexes().CreateOne", idxRet)

	return nil

}
