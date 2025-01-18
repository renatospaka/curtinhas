# curtinhas

series of small snippets of Go code about several topics

## 6 Simple Ways to Optimise Golang

### Declare named with return!

declare the named return with the return type, not only declare the return type.

the above code run result is:

```go
2 - Declare named with return!
declareReturnTypeOnly ExecTime is                                              ->  822ns
2025/01/18 17:49:08 onetwo34five
declareReturnNameTypeWithFmt ExecTime is                                       ->  447ns
2025/01/18 17:49:08 onetwo34five
2025/01/18 17:49:08 declareReturnNameTypeWithLog ExecTime is                   ->  527ns
2025/01/18 17:49:08 onetwo34five
```
