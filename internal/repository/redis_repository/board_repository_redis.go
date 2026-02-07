package redisrepository

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"trellocopy/internal/entity"

	"github.com/redis/go-redis/v9"
)

type RedisBoardRepository struct {
	client *redis.Client
	prefix string
}

func NewRedisBoardRepository(client *redis.Client) *RedisBoardRepository {
	return &RedisBoardRepository{
		client: client,
		prefix: "board:",
	}
}

func (r *RedisBoardRepository) Set(board entity.Board) error {
	ctx := context.Background()
	key := r.prefix + strconv.Itoa(board.ID)

	data, err := json.Marshal(board)

	if err != nil {
		return fmt.Errorf("redis marshal error: %w", err)
	}

	return r.client.Set(ctx, key, data, 0).Err()
}

func (r *RedisBoardRepository) GetByID(id int) (*entity.Board, error) {
	key := r.prefix + strconv.Itoa(id)
	ctx := context.Background()
	data, err := r.client.Get(ctx, key).Result()

	if err == redis.Nil {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("Redis Error: %w", err)
	}

	var board entity.Board

	if err := json.Unmarshal([]byte(data), &board); err != nil {
		return nil, fmt.Errorf("Redis Error: %w", err)
	}
	return &board, nil

}

func (r *RedisBoardRepository) Delete(id int) error {
	ctx := context.Background()
	key := r.prefix + strconv.Itoa(id)
	return r.client.Del(ctx, key).Err()

}
