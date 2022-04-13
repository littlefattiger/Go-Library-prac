package main

import "fmt"

//func a() {
//	for i := 1; i <= 10; i++ {
//		fmt.Println("A: ", i)
//	}
//	wg.Done()
//}
//
//var wg sync.WaitGroup
//
//func b() {
//	for i := 1; i <= 10; i++ {
//		fmt.Println("B: ", i)
//	}
//	wg.Done()
//}
//
//func main() {
//
//	runtime.GOMAXPROCS(8)
//	wg.Add(2)
//	go a()
//	go b()
//	wg.Wait()
//}

//func recv(c chan int) {
//	ret := <-c
//	fmt.Println("接收成功", ret)
//}
//
//func main() {
//	ch := make(chan int)
//	//go recv(ch) // 创建一个 goroutine 从通道接收值
//	ch <- 10
//	fmt.Println("发送成功")
//}
func main() {
	ch := make(chan int, 10)
	for i := 1; i <= 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}
