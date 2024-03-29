package redisServer

import (
	"context"
	logs "github.com/danbai225/go-logs"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"goBackend-demo/bootstrap/global"
	"time"
)

var (
	RedisContext = context.Background()
)

func Init() {
	global.Redis = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"), // 没有密码，默认值
		DB:       viper.GetInt("redis.db"),          // DB 2
	})

	_, err := global.Redis.Ping(RedisContext).Result()
	if err != nil {
		logs.Err("Redis:" + err.Error())
	}

}

// Set 设置值 expire 过期时间
func Set(key string, value interface{}, expire time.Duration) error {
	set := global.Redis.Set(RedisContext, key, value, expire)
	return set.Err()
}

// Get 获取值
func Get(key string) (string, error) {
	get := global.Redis.Get(RedisContext, key)
	return get.Val(), get.Err()
}

// GetSet 获取集合
func GetSet(key string) ([]string, error) {
	set := global.Redis.SMembers(RedisContext, key)
	return set.Val(), set.Err()
}

// PutSet 添加到集合
func PutSet(key string, members []string) error {
	arr := make([]interface{}, 0)
	for _, member := range members {
		arr = append(arr, member)
	}
	setAdd := global.Redis.SAdd(RedisContext, key, arr...)
	return setAdd.Err()
}

// InSet 判断是否在集合中
func InSet(key, val string) (bool, error) {
	set := global.Redis.SIsMember(RedisContext, key, val)
	return set.Val(), set.Err()
}
