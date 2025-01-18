# curtinhas

series of small snippets of Go code about several topics

## 6 Simple Ways to Optimise Golang

### Declare slice with length!

declare your slice with length if possible.

the above code run result is:

```go
3 - Declare slice with length!
2025/01/18 17:49:08 declareReturnSliceNoLength ExecTime is                     -> 10.558µs
2025/01/18 17:49:08 500
2025/01/18 17:49:08 declareReturnSliceWithLength1 ExecTime is                  ->  514ns
2025/01/18 17:49:08 500
```
