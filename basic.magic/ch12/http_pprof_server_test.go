package main

import (
  "testing"
)

func Benchmark_logicOne(b *testing.B) {
    for i := 0; i < b.N; i++ {
        logicOne(10000)
    }
}