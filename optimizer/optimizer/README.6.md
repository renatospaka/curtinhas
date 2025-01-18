# curtinhas

series of small snippets of Go code about several topics

## 6 Simple Ways to Optimise Golang

### Don’t use an empty interface in a map!

map, in Golang, is an essential type and is used almost all the time. 

the above code run result is:

```go
6 - Don’t use an empty interface in a map!
2025/01/18 17:49:08 defineMapFirstWay ExecTime is                    ->  634ns
2025/01/18 17:49:08 defineMapSecWay ExecTime is                      ->  245ns
2025/01/18 17:49:08 defineMapThirdWay ExecTime is                    ->  236ns
2025/01/18 17:49:08 defineMapFourthWay ExecTime is                   ->  203ns
```
