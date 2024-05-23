package storage

type StorageInterface interface {
	Get(partition, key string) ([]byte, error)
	Save(partition, key string, value []byte) error
	Delete(partition, key string) error
}
