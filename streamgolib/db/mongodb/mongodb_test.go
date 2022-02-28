package mongodb

import (
	"context"
	"testing"
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
