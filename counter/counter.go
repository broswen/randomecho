package counter

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"strconv"
	"time"
)

const CounterKey = "COUNTER"

type Service struct {
	rdb *redis.Client
}

func New(address string) (*Service, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       0,
	})
	_, err := client.Ping(context.Background()).Result()
	count := 0
	for err != nil && count < 3 {
		log.Println("error pinging redis, retrying...")
		time.Sleep(3 * time.Second)
		_, err = client.Ping(context.Background()).Result()
		count++
	}

	if err != nil {
		return &Service{}, err
	}
	return &Service{
		rdb: client,
	}, nil
}

func (s *Service) Get(ctx context.Context) (int64, error) {
	val, err := s.rdb.Get(ctx, CounterKey).Result()
	if err != nil {
		if err == redis.Nil {
			return 0, nil
		}
		return 0, err
	}
	count, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (s *Service) Incr(ctx context.Context) (int64, error) {
	val, err := s.rdb.Incr(ctx, CounterKey).Result()
	return val, err
}
