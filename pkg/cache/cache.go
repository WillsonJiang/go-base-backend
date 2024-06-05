package cache

import (
	"context"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

type Cache struct {
	RdbClient *redis.Client
	DB        int
}

func (ca *Cache) ListenExpire(fn func(string)) {
	ctx := context.Background()
	sub := ca.RdbClient.Subscribe(ctx, "__keyevent@"+strconv.Itoa(ca.DB)+"__:expired")
	defer sub.Close()
	for {
		msg := <-sub.Channel()
		payload := msg.Payload
		fn(payload)
	}
}

func (ca *Cache) Set(key string, value any, timeout time.Duration) error {
	ctx := context.Background()
	if err := ca.RdbClient.Set(ctx, key, value, timeout).Err(); err != nil {
		return err
	}
	return nil
}

func (ca *Cache) SetExpires(key string, timeout time.Duration) error {
	ctx := context.Background()
	if err := ca.RdbClient.Expire(ctx, key, timeout).Err(); err != nil {
		return err
	}
	return nil
}

func (ca *Cache) Get(key string) (string, error) {
	ctx := context.Background()
	value, err := ca.RdbClient.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}

func (ca *Cache) GetUint(key string) (uint, error) {
	ctx := context.Background()
	value, err := ca.RdbClient.Get(ctx, key).Uint64()
	if err != nil {
		return 0, err
	}
	return uint(value), nil
}

func (ca *Cache) Del(keys []string) error {
	ctx := context.Background()
	// 若有需要加前綴
	// for index, key := range keys {
	// 	keys[index] = key
	// }
	if err := ca.RdbClient.Del(ctx, keys...).Err(); err != nil {
		return err
	}
	return nil
}

func (ca *Cache) Exists(key string) (bool, error) {
	ctx := context.Background()
	value, err := ca.RdbClient.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return value > 0, nil
}

func (ca *Cache) HSet(hashKey string, data map[string]any) error {
	ctx := context.Background()
	if err := ca.RdbClient.HSet(ctx, hashKey, data).Err(); err != nil {
		return err
	}
	return nil
}

func (ca *Cache) HDel(hashKey string, keys []string) error {
	ctx := context.Background()
	if err := ca.RdbClient.HDel(ctx, hashKey, keys...).Err(); err != nil {
		return err
	}
	return nil
}

func (ca *Cache) HExists(hashKey, key string) (bool, error) {
	ctx := context.Background()
	return ca.RdbClient.HExists(ctx, hashKey, key).Result()
}
