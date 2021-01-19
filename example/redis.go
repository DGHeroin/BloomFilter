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
    bs := BloomFilter.NewRedisBitSet("online_users", conn)
    filter := BloomFilter.NewFilter(bs)

    filter.Add("QAQ")
    test := func(key string) {
        log.Println(key, filter.Exists(key))
    }
    test("QAQ")
    test("PS")
    log.Println("==========")
    filter.DelString("QAQ")
    test("QAQ")
    test("PS")
}
