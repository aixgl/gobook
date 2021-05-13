package main

var balance int
// 存款
func Deposit(amount int) { 
    balance = balance + amount 
}
// 余额
func Balance() int { 
    return balance 
}