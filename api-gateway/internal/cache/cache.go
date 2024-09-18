package cache

import (
	rds "api-gateway/internal/pkg/redis"
	"github.com/redis/go-redis/v9"
	"api-gateway/logger"
	"context"
	"fmt"
	"time"
)

type RedisMethod struct {
	RDS rds.Redis
}

func NewRedis(redis *rds.Redis) *RedisMethod {
	return &RedisMethod{*redis}
}

func (r *RedisMethod) HoldOnUserID(userID string) error {

	expirationTime := time.Now().Add(24 * time.Hour)
	duration := time.Until(expirationTime)
	redisDataName := fmt.Sprintf("%sUser", userID)

	logger.Info("Redis data name and duration", map[string]interface{}{
		"data_name": redisDataName,
		"duration":  duration,
	})

	err := r.RDS.Client.Set(context.Background(), redisDataName, userID, duration).Err()
	if err != nil {
		logger.Error(err)
		return err
	}

	logger.Info("Set UserID to Redis is successfully")
	return nil
}

func (r *RedisMethod) GetUserID(userID string) (string, error) {

	redisDataName := fmt.Sprintf("%sUser", userID)

	logger.Info("Redis data name for get operation", map[string]interface{}{
		"data_name": redisDataName,
	})

	val, err := r.RDS.Client.Get(context.Background(), redisDataName).Result()
	if err == redis.Nil {
		logger.Info("UserID not found in Redis")
		return "", nil 
	} else if err != nil {
		logger.Error(err)
		return "", err
	}

	logger.Info("Retrieved UserID from Redis successfully")
	return val, nil
}
