package main

import (
    "github.com/DGHeroin/BloomFilter"
    "log"
)

func main()  {
    bs := BloomFilter.NewMemoryBitSet()
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
