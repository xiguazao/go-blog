package gredis

import (
	"encoding/json"
	"time"

	//"time"

	"github.com/go-redis/redis"
	//"github.com/gomodule/redigo/redis"

	//"github.com/EDDYCJY/go-gin-example/pkg/setting"
)

var RedisConn *redis.ClusterClient

// Setup Initialize the Redis instance
func Setup() error {
	RedisConn = redis.NewClusterClient(&redis.ClusterOptions {
		Addrs: []string{"10.200.2.170:6379", "10.200.2.171:6379", "10.200.2.172:6379"},
		Password: "",
	})
	return nil
}

// Set a key/value
func Set(key string, data interface{}, time time.Duration) error {
	//conn := RedisConn.Get()
	//defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = RedisConn.Set(key, value, time).Result()
	if err != nil {
		return err
	}
	return nil
}

// Exists check a key
func Exists(key string) bool {
	reply := RedisConn.Exists(key)
	if reply.Val() != 0 {
		return true
	}
	return false
}

// Get get a key
func Get(key string) ([]byte, error) {

	reply, err := RedisConn.Get(key).Result()
	if err != nil {
		return nil, err
	}

	return []byte(reply), nil
}

// Delete delete a kye
func Delete(key string) (int64, error) {

	reply, err := RedisConn.Del(key).Result()
	return reply, err
}

//// LikeDeletes batch delete
//func LikeDeletes(key string) error {
//	conn := RedisConn.Get()
//	defer conn.Close()
//
//	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
//	if err != nil {
//		return err
//	}
//
//	for _, key := range keys {
//		_, err = Delete(key)
//		if err != nil {
//			return err
//		}
//	}
//
//	return nil
//}
