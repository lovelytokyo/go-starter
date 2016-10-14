# 準備

```
go get github.com/kavu/go-resque
go get github.com/simonz05/godis/redis
go get github.com/benmanns/goworker
```

# redis起動

```
redis-server /usr/local/etc/redis.conf
```

# redis接続

```
redis-cli
```

# worker.go実行

```
go run worker.go --queues=default
    ***
From default, [a123456 z96 http://banner-mtb.dspcdn.com/mtbimg/video/bb5/bb59adddc40801a2f9fa10f2116d4185c56a0213]
```

# enqueue.go実行

```
go run enqueue.go
```