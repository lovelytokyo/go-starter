package main

import (
        "github.com/kavu/go-resque" // Import this package
        _ "github.com/kavu/go-resque/godis" // Use godis driver
        "github.com/simonz05/godis/redis" // Import godis package
)

func main() {
        client := redis.New("tcp:127.0.0.1:6379", 0, "") // Godis redis client
        enqueuer := resque.NewRedisEnqueuer("godis", client)

        enqueuer.Enqueue("default", "MyClass", 1, 2, "woot")
}
