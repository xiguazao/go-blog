package main

import (
	//"time"
	"fmt"
	//"github.com/gomodule/redigo/redis"
	"github.com/go-redis/redis"

	//"github.com/EDDYCJY/go-gin-example/pkg/setting"
)


func main(){
	RedisConn := redis.NewClusterClient(&redis.ClusterOptions {
		Addrs: []string{"10.200.2.170:6379", "10.200.2.171:6379", "10.200.2.172:6379"},
		Password: "",
	})
	a := RedisConn.Get("name")
	fmt.Print(a)
	res, err := RedisConn.Get("name").Result()
	fmt.Print(res)
	fmt.Print(err)
	//b := RedisConn.Set("name", "test", 0)
	//fmt.Print(b)
	//d := RedisConn.Set("name1", "test2", 0)
	//fmt.Print(d)

	c := RedisConn.Get("name1")
	fmt.Print(c)
}
