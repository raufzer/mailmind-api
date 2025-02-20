package config

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RedisConfig struct {
	Client *redis.Client
}

func ConnectRedis(config *AppConfig) *RedisConfig {

	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisURI,
		Password: config.RedisPassword,
		DB:       0,
	})

	ctx := context.Background()

	if err := client.Ping(ctx).Err(); err != nil {
		log.Fatalf("Error while connecting to Redis: %v", err)
	}

	log.Println("Connected to Redis successfully")

	return &RedisConfig{
		Client: client,
	}
}

var DB *mongo.Database

func ConnectDB() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017") 
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")

	DB = client.Database("mailmind") // Set your database name
}

func GetCollection(collectionName string) *mongo.Collection {
	return DB.Collection(collectionName)
}
