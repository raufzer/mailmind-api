package interfaces

import (
	"context"
	"mailmind-api/pkg/utils"
	"time"
)

type RedisRepository interface {
	StoreAssetCache(ctx context.Context, assetID string, assetType string, data *utils.AssetCache, expiry time.Duration) error
	GetAssetCache(ctx context.Context, assetID string, assetType string) (*utils.AssetCache, error)
	InvalidateAssetCache(ctx context.Context, assetID string, assetType string) error
}
