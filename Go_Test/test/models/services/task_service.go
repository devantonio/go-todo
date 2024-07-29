package services

import (
    "context"
    "log"

    "appname/internal/models"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

type TaskService struct {
    collection *mongo.Collection
}

func NewTaskService(client *mongo.Client) *TaskService {
    return &TaskService{
        collection: client.Database("your_database").Collection("tasks"),
    }
}

func (s *TaskService) GetTasks() ([]models.Task, error) {
    var tasks []models.Task
    cursor, err := s.collection.Find(context.Background(), bson.D{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.Background())

    for cursor.Next(context.Background()) {
        var task models.Task
        err := cursor.Decode(&task)
        if err != nil {
            return nil, err
        }
        tasks = append(tasks, task)
    }

    return tasks, nil
}

func (s *TaskService) CreateTask(task models.Task) (*mongo.InsertOneResult, error) {
    return s.collection.InsertOne(context.Background(), task)
}

func (s *TaskService) GetTaskByID(id primitive.ObjectID) (models.Task, error) {
    var task models.Task
    err := s.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&task)
    return task, err
}

func (s *TaskService) UpdateTask(id primitive.ObjectID, task models.Task) (*mongo.UpdateResult, error) {
    return s.collection.UpdateOne(
        context.Background(),
        bson.M{"_id": id},
        bson.D{
            {"$set", bson.D{
                {"title", task.Title},
                {"completed", task.Completed},
            }},
        },
    )
}

func (s *TaskService) DeleteTask(id primitive.ObjectID) (*mongo.DeleteResult, error) {
    return s.collection.DeleteOne(context.Background(), bson.M{"_id": id})
}
