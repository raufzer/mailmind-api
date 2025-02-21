package interfaces

import (
	"context"
	"mailmind-api/pkg/utils"
	"time"
)

type RedisRepository interface {
	StoreRefreshToken(ctx context.Context, UserID, refreshToken string, expiry time.Duration) error
	GetRefreshToken(ctx context.Context, UserID string) (string, error)
	InvalidateRefreshToken(ctx context.Context, UserID string) error
	StoreAssetCache(ctx context.Context, assetID string, assetType string, data *utils.AssetCache, expiry time.Duration) error
	GetAssetCache(ctx context.Context, assetID string, assetType string) (*utils.AssetCache, error)
	InvalidateAssetCache(ctx context.Context, assetID string, assetType string) error
}
