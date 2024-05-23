package storage

import "context"

type StorageInterface interface {
	Get(ctx context.Context, partition, key string) ([]byte, error)
	Save(ctx context.Context, partition, key string, value []byte) error
	Delete(ctx context.Context, partition, key string) error
}
