package session

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis"
)

var SessionLocal = session.New(session.Config{
	Storage: redis.New(redis.Config{}),
})

func Set(key string, val []byte, exp time.Duration) error {
	return SessionLocal.Storage.Set(key, val, exp)
}

func Get(key string) ([]byte, error) {
	return SessionLocal.Storage.Get(key)
}
