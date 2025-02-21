package config

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
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
type DatabaseConfig struct {
	Client         *mongo.Client
	UserCollection *mongo.Collection
	Ctx            context.Context
}

func ConnectDatabase(config *AppConfig) *DatabaseConfig {
	ctx := context.TODO()

	clientOptions := options.Client().ApplyURI(config.DatabaseURI)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Error while connecting to MongoDB: ", err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("Error while pinging MongoDB: ", err)
	}

	log.Println("Connected to MongoDB successfully")

	return &DatabaseConfig{
		Client:         client,
		UserCollection: client.Database("userdb").Collection("users"),
		Ctx:            ctx,
	}
}