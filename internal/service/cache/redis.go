package cache

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

// Redis ...
type Redis struct {
	pool *redis.Pool
}

// New ...
func New() (*Redis, error) {
	pool := &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 60 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "cont_redis")
		},
	}

	return &Redis{pool: pool}, nil
}

// Set ...
func (r *Redis) Set(key, value string, timeout time.Duration) error {
	conn := r.pool.Get()
	defer conn.Close()

	return nil
}

// Get ...
func (r *Redis) Get(key string) (string, error) {
	return "", nil
}
