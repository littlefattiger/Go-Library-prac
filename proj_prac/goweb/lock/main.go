package main

import (
	"fmt"
	"strconv"
	"sync"
)

//var (
//	x int64
//
//	wg sync.WaitGroup // 等待组
//
//	m sync.Mutex
//)
//
//// add 对全局变量x执行5000次加1操作
//func add() {
//	for i := 0; i < 5000; i++ {
//		m.Lock() // 修改x前加锁
//		x = x + 1
//		m.Unlock() // 改完解锁
//	}
//	wg.Done()
//}
//
//func main() {
//	wg.Add(2)
//
//	go add()
//	go add()
//
//	wg.Wait()
//	fmt.Println(x)
//}
//var (
//	x       int64
//	wg      sync.WaitGroup
//	mutex   sync.Mutex
//	rwMutex sync.RWMutex
//)
//
//// writeWithLock 使用互斥锁的写操作
//func writeWithLock() {
//	mutex.Lock() // 加互斥锁
//	x = x + 1
//	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
//	mutex.Unlock()                    // 解互斥锁
//	wg.Done()
//}
//
//// readWithLock 使用互斥锁的读操作
//func readWithLock() {
//	mutex.Lock()                 // 加互斥锁
//	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
//	mutex.Unlock()               // 释放互斥锁
//	wg.Done()
//}
//
//// writeWithLock 使用读写互斥锁的写操作
//func writeWithRWLock() {
//	rwMutex.Lock() // 加写锁
//	x = x + 1
//	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
//	rwMutex.Unlock()                  // 释放写锁
//	wg.Done()
//}
//
//// readWithRWLock 使用读写互斥锁的读操作
//func readWithRWLock() {
//	rwMutex.RLock()              // 加读锁
//	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
//	rwMutex.RUnlock()            // 释放读锁
//	wg.Done()
//}
//
//func do(wf, rf func(), wc, rc int) {
//	start := time.Now()
//	// wc个并发写操作
//	for i := 0; i < wc; i++ {
//		wg.Add(1)
//		go wf()
//	}
//
//	//  rc个并发读操作
//	for i := 0; i < rc; i++ {
//		wg.Add(1)
//		go rf()
//	}
//
//	wg.Wait()
//	cost := time.Since(start)
//	fmt.Printf("x:%v cost:%v\n", x, cost)
//
//}
//
//func main() {
//	do(writeWithLock, readWithLock, 10, 1000)
//
//	do(writeWithRWLock, readWithRWLock, 10, 1000)
//}

var m = make(map[string]int)

func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			set(key, n)
			fmt.Printf("k=:%v,v:=%v\n", key, get(key))
			wg.Done()
		}(i)
	}
	wg.Wait()
}
