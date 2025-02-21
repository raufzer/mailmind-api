package interfaces

import (
	"mailmind-api/pkg/utils"
	"time"
)

type RedisRepository interface {
	StoreResetToken(email, token string, expiry time.Duration) error
	GetResetToken(email string) (string, error)
	InvalidateResetToken(email string) error
	StoreRefreshToken(UserID, refreshToken string, expiry time.Duration) error
	GetRefreshToken(UserID string) (string, error)
	InvalidateRefreshToken(UserID string) error
	StoreAssetCache(assetID string, assetType string, data *utils.AssetCache, expiry time.Duration) error
	GetAssetCache(assetID string, assetType string) (*utils.AssetCache, error)
	InvalidateAssetCache(assetID string, assetType string) error
}
