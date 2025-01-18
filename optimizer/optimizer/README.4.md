# curtinhas

series of small snippets of Go code about several topics

## 6 Simple Ways to Optimise Golang

### Join string using strings.Builder!

Jjin string using strings.Builder, if possible!

the above code run result is:

```go
4 - Join string using strings.Builder!
2025/01/18 17:49:08 joinStrDirect ExecTime is                     ->  791ns
2025/01/18 17:49:08 6 simple ways to optimise Golang - this is the catch
2025/01/18 17:49:08 joinStrWithSprintf ExecTime is                -> 1.732µs
2025/01/18 17:49:08 6 simple ways to optimise Golang - this is the catch
2025/01/18 17:49:08 joinStrWithBuilder ExecTime is                ->  489ns
2025/01/18 17:49:08 6 simple ways to optimise Golang - this is the catch
```
