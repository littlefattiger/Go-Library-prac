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
