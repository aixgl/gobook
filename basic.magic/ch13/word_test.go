package main

import "testing"
import "github.com/stretchr/testify/assert"

// ------------------------go 自带-------------------------------
func TestHasChinese(t *testing.T) {
    s1 := "I am a 科学家"
    if !hasChinese(s1) {
        t.Errorf("hasChinese(%s) == false; expected:true", s1)
    }
    s2 := "Who are you ?"
    if hasChinese(s2) {
        t.Errorf("hasChinese(%s) == true; expected:false", s2)
    }
    t.Log("---", t.Name(), "END ----")
}

func TestAbs(t *testing.T) {
    if abs(-2) != 2 {
        t.Errorf("abs(%d) check error", -2)
    }
    if abs(2) !=2 {
        t.Errorf("abs(%d) check error", 2)
    }
}

func TestDivide1(t *testing.T) {
    a, b := 3, 0
    if b == 0 {
        //t.Fatalf("b == 0") 无参数可用t.Fail()
        t.Fatal()
    }
    divide(a, b)
}
/* 取消注释 test则会阻断
func TestDivide2(t *testing.T) {
    a, b := 3, 0
    divide(a, b)
}
*/

func TestDivide3(t *testing.T) {
    t.Fail()
    t.Log("---", t.Name(), "---")
    a, b := 3, 1
    divide(a, b)
}
func TestDivide4(t *testing.T) {
    t.FailNow()
    t.Log("---", t.Name(), "---")
    a, b := 3, 1
    divide(a, b)
}

// ------------------------go 第三方测试框架-------------------------------
func TestAbsThird(t *testing.T) {
    assert.Equal(t, 2, abs(-2))
    assert.Equal(t, 2, abs(2))
}
