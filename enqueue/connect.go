package main

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"flag"
	"fmt"
	"math/big"
	"runtime"
	"strconv"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/mr51m0n/gorc"
)

var (
	pool *redis.Pool
	host string = "localhost:6379" // redis-server接続情報
	key int = 150000 // 投入するKEYの数
	val int = 10000 // 文字列の長さ
)

// gorc
var gorc0 gorc.Gorc

func init() {
	gorc0.Init()
}

// メイン処理
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	flag.Parse()

	// Timer Set
	start := time.Now()

	// redis Pooling Start
	pool = newPool(host)

	// Redis DataInsert
	for i := 0; i < key; i++ {
		gorc0.Inc()
		go setredis(i, val)

		gorc0.WaitLow(100)
	}


	// Log Print
	fmt.Print(key)
	fmt.Printf("Keys - Set Done.\n")

	// Timer Print
	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
}

// redis ConnectionPooling
func newPool(server string) *redis.Pool {
	return &redis.Pool{

		MaxIdle: 3,
		IdleTimeout: 240 * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)

			if err != nil {
				return nil, err
			}
			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

// redisにデータを投入するだけの処理
func setredis(k, v int) {

	c := pool.Get()
	defer c.Close()
	defer gorc0.Dec()


	vals := random(v)
	keys := GetMD5Hash(vals)


	c.Do("SET", keys, vals)


	// 100処理ごとにログ表示(非同期なので順番が適当)
	if k%100 == 0 {
		fmt.Print(k)
		fmt.Printf("Keys - Set Done.\n")
	}
}

// stringからmd5から作る
func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// val用にrandomな文字列を生成する
func random(length int) string {
	const base = 36
	size := big.NewInt(base)
	n := make([]byte, length)
	for i, _ := range n {
		c, _ := rand.Int(rand.Reader, size)
		n[i] = strconv.FormatInt(c.Int64(), base)[0]
	}
	return string(n)
}
