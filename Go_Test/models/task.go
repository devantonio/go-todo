package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Task represents a task entity in the application.
type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title"`
	Completed   bool               `bson:"completed"`
	Description string             `bson:"description"`
}

type FeaturedComment struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	MediaType        string             `bson:"media_type"`
	UserID           primitive.ObjectID `bson:"user_id,omitempty"`
	Name             string             `bson:"name"`
	Username         string             `bson:"username"`
	Comment          string             `bson:"comment"`
	CommentLikeCount int                `bson:"commentLikeCount"`
	ReplyCount       int                `bson:"replyCount"`
	MediaID          string             `bson:"media_id"`
	Created          string             `bson:"created"`
}
