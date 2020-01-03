package cache

import (
	"errors"
	"time"

	"github.com/garyburd/redigo/redis"
)

const (
	// DefaultSessionTTL explicitly expires the key after 10 minutes
	DefaultSessionTTL time.Duration = (10 * time.Minute)
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

	s := set(conn, timeout)

	if _, err := s(key, value); err != nil {
		return err
	}

	return nil
}

// Get ...
func (r *Redis) Get(key string) (string, error) {
	conn := r.pool.Get()
	defer conn.Close()

	reply, err := conn.Do("GET", key)
	if err != nil {
		return "", err
	}

	value, ok := reply.(string)
	if !ok {
		return "", errors.New("reply is not string type")
	}

	return value, nil
}

type setter func(key, value string) (interface{}, error)

func set(conn redis.Conn, timeout time.Duration) setter {
	return func(key, value string) (interface{}, error) {
		if timeout != 0 {
			return conn.Do("SETEX", key, value, timeout.Seconds())
		}

		return conn.Do("SET", key, value)
	}
}
