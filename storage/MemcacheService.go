package storage

import (
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
)

type MemcachedImp struct {
	client *memcache.Client
}

func NewStorageMemcached() StorageInterface {
	memcachedClient := memcache.New("localhost:11211")
	return &MemcachedImp{
		memcachedClient,
	}
}

func (s *MemcachedImp) Get(partition, key string) ([]byte, error) {
	get, err := s.client.Get(fmt.Sprintf("%s-%s", partition, key))
	if err != nil {
		return nil, err
	}
	return get.Value, nil
}

func (s *MemcachedImp) Save(partition, key string, value []byte) error {
	err := s.client.Add(
		&memcache.Item{
			Key:   fmt.Sprintf("%s-%s", partition, key),
			Value: value,
		})
	if err != nil {
		return err
	}
	return nil
}

func (s *MemcachedImp) Delete(partition, key string) error {
	err := s.client.Delete(fmt.Sprintf("%s-%s", partition, key))
	if err != nil {
		return err
	}
	return nil
}
