package storage

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type RedisImp struct {
	client *redis.Client
}

func NewStorageRedis() StorageInterface {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	return &RedisImp{
		redisClient,
	}
}

func (s *RedisImp) Get(ctx context.Context, partition, key string) ([]byte, error) {
	std := s.client.Get(
		ctx,
		fmt.Sprintf("%s-%s", partition, key),
	)
	if std.Err() != nil {
		return nil, std.Err()
	}
	result, err := std.Result()
	if err != nil {
		return nil, err
	}
	return []byte(result), nil
}

func (s *RedisImp) Save(ctx context.Context, partition, key string, value []byte) error {
	std := s.client.Set(
		ctx,
		fmt.Sprintf("%s-%s", partition, key),
		value,
		0,
	)
	if std.Err() != nil {
		return std.Err()
	}
	return nil
}

func (s *RedisImp) Delete(ctx context.Context, partition, key string) error {
	rsp := s.client.Del(ctx, fmt.Sprintf("%s-%s", partition, key))
	if rsp.Err() != nil {
		return rsp.Err()
	}
	return nil
}
