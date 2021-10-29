package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

var (
	database   string
	collection string
)

const (
	// environment variables
	mongoDBConnectionStringEnvVarName = "MONGODB_CONNECTION_STRING"
	mongoDBDatabaseEnvVarName         = "MONGODB_DATABASE"
	mongoDBCollectionEnvVarName       = "MONGODB_COLLECTION"
)

// connects to MongoDB
func connect() *mongo.Client {
	mongoDBConnectionString := os.Getenv(mongoDBConnectionStringEnvVarName)
	if mongoDBConnectionString == "" {
		log.Fatal("missing environment variable: ", mongoDBConnectionStringEnvVarName)
	}

	database = os.Getenv(mongoDBDatabaseEnvVarName)
	if database == "" {
		log.Fatal("missing environment variable: ", mongoDBDatabaseEnvVarName)
	}

	collection = os.Getenv(mongoDBCollectionEnvVarName)
	if collection == "" {
		log.Fatal("missing environment variable: ", mongoDBCollectionEnvVarName)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	clientOptions := options.Client().ApplyURI(mongoDBConnectionString).SetDirect(true)
	c, err := mongo.NewClient(clientOptions)

	err = c.Connect(ctx)

	if err != nil {
		log.Fatalf("unable to initialize connection %v", err)
	}
	err = c.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("unable to connect %v", err)
	}
	return c
}

// Create a user
func Create(visited string, id primitive.ObjectID) {
	c := connect()
	ctx := context.Background()
	defer c.Disconnect(ctx)

	userCollection := c.Database(database).Collection(collection)
	r, err := userCollection.InsertOne(ctx, User{
		ID:      id,
		Visited: []string{visited},
	})
	if err != nil {
		log.Fatalf("failed to add user %v", err)
	}
	fmt.Println("added user", r.InsertedID)
}

// AddVisited adds a visited to existing user
func AddVisited(userID string, newVisited string) {
	c := connect()
	ctx := context.Background()
	defer c.Disconnect(ctx)

	userCollection := c.Database(database).Collection(collection)
	println("trying to parse objectid: ", userID)
	oid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Fatalf("failed to update user %v", err)
	}
	filter := bson.D{{"_id", oid}}
	update := bson.M{"$push": bson.M{"visited": newVisited}}
	_, err = userCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatalf("failed to update user %v", err)
	}
}

type User struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Visited []string           `bson:"visited"`
}
