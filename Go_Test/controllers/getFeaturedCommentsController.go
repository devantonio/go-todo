package controllers

import (
	"context"
	"fmt"
	"go_test/models"
	"time"

	//"go_test/config"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	// Add this import
)

func getFeaturedComments() ([]models.FeaturedComment, error) {
	//cfg := config.LoadConfig()
	var FeaturedComments []models.FeaturedComment

	client, err := DB()
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(context.Background())

	// Use the client to access the database and collection
	collection := client.Database("SR_DB").Collection("comments")

	// Aggregation pipeline
	pipeline := mongo.Pipeline{
		{{
			Key: "$addFields", Value: bson.D{{
				Key: "combinedCount", Value: bson.D{{
					Key: "$add", Value: bson.A{"$commentLikeCount", "$replyCount"},
				}},
			}},
		}},
		{{Key: "$sort", Value: bson.D{{Key: "combinedCount", Value: -1}}}},
		{{
			Key: "$group", Value: bson.D{
				{Key: "_id", Value: "$combinedCount"},
				{Key: "docs", Value: bson.D{{Key: "$push", Value: "$$ROOT"}}},
			},
		}},
		{{Key: "$sort", Value: bson.D{{Key: "_id", Value: -1}}}},
		{{Key: "$limit", Value: 1}},
		{{Key: "$unwind", Value: "$docs"}},
		{{Key: "$replaceRoot", Value: bson.D{{Key: "newRoot", Value: "$docs"}}}},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &FeaturedComments); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Comments with the highest combined count:")
	for _, FeaturedComment := range FeaturedComments {
		cursor.Decode(&FeaturedComment)
		FeaturedComments = append(FeaturedComments, FeaturedComment)
	}

	// for cursor.Next(context.Background()) {
	// 	var FeaturedComment models.FeaturedComment
	// 	cursor.Decode(&FeaturedComment)
	// 	FeaturedComments = append(FeaturedComments, FeaturedComment)
	// }

	fmt.Println(FeaturedComments)

	return FeaturedComments, nil
}
