-   Ad hoc confinement
-   Lexical confinement -> This is preferred.
for-select Loop
```
for { // Either loop infinitely or range over something
 select {
 // Do some work with channels
 }
}

for {
 select {
 case <-done:
 return
 default:
 // Do non-preemptable work
 }
}
```
Using the signal
```
newRandStream := func(done <-chan interface{}) <-chan int {
 randStream := make(chan int)
 go func() {
 defer fmt.Println("newRandStream closure exited.")
 defer close(randStream)
 for {
 select {
 case randStream <- rand.Int():
 case <-done:
 return
 }
 }
 }()
return randStream
}
done := make(chan interface{})
randStream := newRandStream(done)
fmt.Println("3 random ints:")
for i := 1; i <= 3; i++ {
 fmt.Printf("%d: %d\n", i, <-randStream)
}
close(done)
// Simulate ongoing work
time.Sleep(1 * time.Second)

```
-   or-channel
-   Error Handling
-   pipeline stage -> higher order functions and monads
    - If a stage is blocked on retrieving a value from the incoming channel, it will become unblocked when that channel is closed. 
    - Fan-out is a term to describe the process of starting multiple goroutines to handle input from the pipeline, and fan-in is a term to describe the process of combining multiple results into one channel


