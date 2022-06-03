# Struct in better order!

## In golang, data type and size list is following:

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


## The above code run result is:
Size of main.Employee1 struct: 216 bytes
Size of main.Employee2 struct: 192 bytes
Size of main.Employee3 struct: 208 bytes