package cache

type CacheInterface interface {
	Iterator() Iterator
	Get(key []byte) ([]byte, error)
	Set(key []byte, v []byte, expire int) error
	Del(key []byte)

	EntryCount() (entryCount int64)
	HitCount() int64
	MissCount() int64
}
