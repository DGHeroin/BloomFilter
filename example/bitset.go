package main

import (
    "github.com/DGHeroin/BloomFilter"
    "log"
)

func main()  {
    bs := BloomFilter.NewMemoryBitSet()
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
