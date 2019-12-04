package main

import (
	"fmt"
	"time"
)

// 并发编程

func main() {
	for i:=0; i<1000; i++ {
		// go 关键字定义并发编程函数，可以定义匿名函数，也可以调用已定义好的函数
		// 写法1，go func() {} ()
		// 写法2，go hello() {}，前提是一定义func hello(){}这个方法
		// 以下实现以并发的方式调用匿名函数func
		go func() { // 表示并发执行func函数，而函数中将一直执行输出打印操作
			for { // for表示一直执行
				fmt.Printf("hello from goroutine %d\n", i)
			}
		}()
	}
	// main和启动的goroutine同时执行，但是main先结束，因此整个程序都将结束，而goroutine还没有开始执行就结束了
	// 所以添加一个毫秒的延时，让goroutine先执行完，再执行main
	time.Sleep(time.Microsecond)
}
