package redis

import (
	"context"
	"fmt"
	"github.com/spf13/viper"

	"github.com/go-redis/redis/v8"
)

var (
	client *redis.Client
	Nil    = redis.Nil
	ctx    = context.Background()
)

// Init 初始化连接
func Init() (err error) {
	client = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetString("redis.port")),
		Password:     viper.GetString("redis.password"), // no password set
		DB:           viper.GetInt("redis.db"),          // use default DB
		PoolSize:     viper.GetInt("redis.pool_size"),
		MinIdleConns: 10,
	})

	_, err = client.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}

func Close() {
	_ = client.Close()
}
