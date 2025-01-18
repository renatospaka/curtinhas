# curtinhas

series of small snippets of Go code about several topics

## 6 Simple Ways to Optimise Golang

### Struct in better order!

in golang, data type and size list is following:

```go
Data Type        Size
bool            1 byte
int16           2 bytes
int32           4 bytes
int64           8 bytes
int             8 bytes
string          16 bytes
float32         4 bytes
float64         8 bytes
uint32          4 bytes
uint64          8 bytes
nil interface{}  16 bytes
time.Time        24 bytes   // it is a struct, so its size is unstable.
time.Timer       80 bytes   // it is a struct, so its size is unstable.
time.Duration    8 bytes
[]byte           24 bytes
```

The above code run result is:

```go
1 - Struct in better order!
Size of optimizer.employee1 struct: 152 bytes
Size of optimizer.employee2 struct: 128 bytes
Size of optimizer.employee3 struct: 144 bytes
```
