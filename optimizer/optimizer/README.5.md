# curtinhas

series of small snippets of Go code about several topics

## 6 Simple Ways to Optimise Golang

### Instead for range with for loop!

don’t use for range, use for loop instead, if possible.

the above code run result is:

```go
5 - Instead for range with for loop!
2025/01/18 17:49:08 joinStrWithSprintf ExecTime is                -> 2.33µs
2025/01/18 17:49:08 6 simple ways to optimise Golang - this is the catch
2025/01/18 17:49:08 joinStrWithBuilder ExecTime is                            -> 10.631µs
2025/01/18 17:49:08 6 simple ways to optimise Golang - this is the catch
2025/01/18 17:49:08 joinStrWithBuilderForLoop ExecTime is                     ->  426ns
2025/01/18 17:49:08 6 simple ways to optimise Golang - this is the catch
```
