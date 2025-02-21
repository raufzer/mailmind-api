package redis

import (
	"context"
	"encoding/json"
	"fmt"
	repositoryInterfaces "mailmind-api/internal/repositories/interfaces"
	"mailmind-api/pkg/utils"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisRepository struct {
	redisClient *redis.Client
}

func NewRedisRepository(redisClient *redis.Client) repositoryInterfaces.RedisRepository {
	return &RedisRepository{
		redisClient: redisClient,
	}
}

func (r *RedisRepository) StoreRefreshToken(ctx context.Context, UserID, refreshToken string, expiry time.Duration) error {
	key := fmt.Sprintf("refresh_token:%s", UserID)
	if err := r.redisClient.Set(context.Background(), key, refreshToken, expiry).Err(); err != nil {
		return fmt.Errorf("redis: failed to store refresh token for UserID %s: %w", UserID, err)
	}
	return nil
}

func (r *RedisRepository) GetRefreshToken(ctx context.Context, UserID string) (string, error) {
	key := fmt.Sprintf("refresh_token:%s", UserID)
	result, err := r.redisClient.Get(context.Background(), key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", redis.Nil
		}
		return "", fmt.Errorf("redis: failed to get refresh token for UserID %s: %w", UserID, err)
	}
	return result, nil
}

func (r *RedisRepository) InvalidateRefreshToken(ctx context.Context, UserID string) error {
	key := fmt.Sprintf("refresh_token:%s", UserID)
	if err := r.redisClient.Del(context.Background(), key).Err(); err != nil {
		return fmt.Errorf("redis: failed to delete refresh token for user_id %s: %w", UserID, err)
	}
	return nil
}

func (r *RedisRepository) StoreAssetCache(ctx context.Context, assetID string, assetType string, data *utils.AssetCache, expiry time.Duration) error {
	key := fmt.Sprintf("asset:%s:%s", assetType, assetID)

	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("redis: failed to marshal asset data: %w", err)
	}

	if err := r.redisClient.Set(context.Background(), key, jsonData, expiry).Err(); err != nil {
		return fmt.Errorf("redis: failed to store asset cache for ID %s: %w", assetID, err)
	}
	return nil
}

func (r *RedisRepository) GetAssetCache(ctx context.Context, assetID string, assetType string) (*utils.AssetCache, error) {
	key := fmt.Sprintf("asset:%s:%s", assetType, assetID)

	result, err := r.redisClient.Get(context.Background(), key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, fmt.Errorf("redis: failed to get asset cache for ID %s: %w", assetID, err)
	}

	var assetCache utils.AssetCache
	if err := json.Unmarshal([]byte(result), &assetCache); err != nil {
		return nil, fmt.Errorf("redis: failed to unmarshal asset data: %w", err)
	}

	return &assetCache, nil
}

func (r *RedisRepository) InvalidateAssetCache(ctx context.Context, assetID string, assetType string) error {
	key := fmt.Sprintf("asset:%s:%s", assetType, assetID)

	if err := r.redisClient.Del(context.Background(), key).Err(); err != nil {
		return fmt.Errorf("redis: failed to invalidate asset cache for ID %s: %w", assetID, err)
	}
	return nil
}
