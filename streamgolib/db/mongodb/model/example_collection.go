package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Podcast represents the schema for the "Podcasts" collection
type Podcast struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Title  string             `bson:"title,omitempty"`
	Author string             `bson:"author,omitempty"`
	Tags   []string           `bson:"tags,omitempty"`
}

// Episode represents the schema for the "Episodes" collection
type Episode struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Podcast     primitive.ObjectID `bson:"podcast,omitempty"`
	Title       string             `bson:"title,omitempty"`
	Description string             `bson:"description,omitempty"`
	Duration    int32              `bson:"duration,omitempty"`
}

/*
id, _ := primitive.ObjectIDFromHex("5e3b37e51c9d4400004117e6")

matchStage := bson.D{{"$match", bson.D{{"podcast", id}}}}
groupStage := bson.D{{"$group", bson.D{{"_id", "$podcast"}, {"total", bson.D{{"$sum", "$duration"}}}}}}

showInfoCursor, err := episodesCollection.Aggregate(ctx, mongo.Pipeline{matchStage, groupStage})
if err != nil {
    panic(err)
}
var showsWithInfo []bson.M
if err = showInfoCursor.All(ctx, &showsWithInfo); err != nil {
    panic(err)
}
fmt.Println(showsWithInfo)


lookupStage := bson.D{{"$lookup", bson.D{{"from", "podcasts"}, {"localField", "podcast"}, {"foreignField", "_id"}, {"as", "podcast"}}}}
unwindStage := bson.D{{"$unwind", bson.D{{"path", "$podcast"}, {"preserveNullAndEmptyArrays", false}}}}

showLoadedCursor, err := episodesCollection.Aggregate(ctx, mongo.Pipeline{lookupStage, unwindStage})
if err != nil {
    panic(err)
}
var showsLoaded []bson.M
if err = showLoadedCursor.All(ctx, &showsLoaded); err != nil {
    panic(err)
}
fmt.Println(showsLoaded)

*/
