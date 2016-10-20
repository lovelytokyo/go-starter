package main

import (
	"github.com/garyburd/redigo/redis"
	"time"
	"encoding/json"
	"fmt"
)

type RedisClient struct {
	Conn redis.Conn
}

type jobArg interface{}

type job struct {
	Class string   `json:"class"`
	Args  []jobArg `json:"args"`
}

var pool *redis.Pool

func main() {
	pool = &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 30,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
		TestOnBorrow: func(rc redis.Conn, t time.Time) error {
			_, err := rc.Do("PING")
			return err
		},
	}

	conn := pool.Get()

	jobs := getJobList()

	NewRedisClient(conn).ListPush("default_test", jobs)
}

func getJobList() []interface{} {

	var jobs []*job
	var list []interface{}

	jobs = append(jobs, &job {
			Class: "MyClass",
			Args: []jobArg{1, 2, "woot"},
		},&job {
			Class: "MyClass",
			Args: []jobArg{3, 4, "boox"},
		})

	for _, job := range jobs {
		jobJson, err := json.Marshal(job)
		list = append(list, string(jobJson))

		if err != nil {
			fmt.Printf("err : %+v", err)
		}
	}

	return list
}

func NewRedisClient (conn redis.Conn) *RedisClient {
	return &RedisClient{Conn: conn}
}

func (r *RedisClient) ListPush(key string, list []interface{}) {
	r.Conn.Do("RPUSH", append([]interface{}{key}, list...)...)
}
