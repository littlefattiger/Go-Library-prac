- Thus, goroutines can be considered a special class of coroutine.
- Fail:
```
var wg sync.WaitGroup
for _, salutation := range []string{"hello", "greetings", "good day"} {
wg.Add(1) go func() {
defer wg.Done()
            fmt.Println(salutation)
        }()
} wg.Wait()
```
Success:
```
var wg sync.WaitGroup
for _, salutation := range []string{"hello", "greetings", "good day"} {
wg.Add(1)
go func(salutation string) {
defer wg.Done()
            fmt.Println(salutation)
        }(salutation)
} wg.Wait()
```
Gorouting is:
1. Lightweight, only use 1k per gorouting
2. It is fast than context switch. Sending/Receiving info from channel is fast.
```
func BenchmarkContextSwitch(b *testing.B) { var wg sync.WaitGroup
begin := make(chan struct{})
c := make(chan struct{})
var token struct{} sender := func() {
defer wg.Done()
<-begin fori:=0;i<b.N;i++{
c <- token }
}
receiver := func() {
defer wg.Done()
<-begin fori:=0;i<b.N;i++{
<-c }
}
wg.Add(2)
go sender() go receiver() b.StartTimer() close(begin) wg.Wait()
}
```
#### The sync Package
concurrency primitives
- WaitGroup is a great way to wait for a set of concurrent operations to complete when you either don’t care about the result of the concurrent operation, or you have other means of collecting their results.
- Mutex and RWMutex
- Cond
- Once ->doing once for our instance, no matter what function we have inside, and how many inside
- Pool
- channal
- select

```
chanOwner := func() <-chan int {
 resultStream := make(chan int, 5)
 go func() {
 defer close(resultStream)
 for i := 0; i <= 5; i++ {
 resultStream <- i
 }
 }()
 return resultStream
}
resultStream := chanOwner()
for result := range resultStream {
 fmt.Printf("Received: %d\n", result)
}
fmt.Println("Done receiving!")
```

GOMAXPROCS

