package main

import (
    "context"
    "github.com/DGHeroin/BloomFilter"
    "github.com/go-redis/redis/v8"
    "log"
)

func main()  {
    client := redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
    })
    conn := client.Conn(context.Background())
    if cmd := conn.Ping(context.Background()); cmd.Err() != nil {
        log.Println(cmd.Err())
        return
    }
    bs := BloomFilter.NewRedisBitSet("online_users", conn)
    filter := BloomFilter.NewFilter(bs)

    filter.Add("111")
    filter.Add("222")
    test := func(key string) {
       log.Println(key, filter.Exists(key))
    }
    test("111")
    test("333")
    log.Println("==========")
    filter.DelString("111")
    test("111")
    test("222")

    log.Println("Count:", filter.Count())
}
