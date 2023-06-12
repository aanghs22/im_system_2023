import {
	"context"
	"time"
	"github.com/redis/go-redis/v9"
	"encoding/json"
}

type RedisClient struct {
	cli *redis.Client
}

func (client *RedisClient) Redis(c context.Context, address, password string) {
	cli := redis.NewClient(&redis.Options{
		Addr: address,
		Password: password,
		DB: 0,
	})

	err := cli.Ping(c).Err();

	if err != nil {
		return err
	}

	c.client = cli
	return nil;
}