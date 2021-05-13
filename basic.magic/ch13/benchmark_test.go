package main

import "testing"

func BenchmarkHasChinese(b *testing.B) {
    s1 := "I am a 科学家"
    for i:=0; i < b.N; i++ {
        hasChinese(s1)
    }
    b.Log("---", b.Name(), "END ----")
}