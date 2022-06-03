# Struct in better order!

## In golang, data type and size list is following:
Data Type        Size <br/>
bool            1 byte <br/>
int16           2 bytes <br/>
int32           4 bytes <br/>
int64           8 bytes <br/>
int             8 bytes <br/>
string          16 bytes <br/>
float32         4 bytes <br/>
float64         8 bytes <br/>
uint32          4 bytes <br/>
uint64          8 bytes <br/>
nil interface{}  16 bytes <br/>
time.Time        24 bytes   // it is a struct, so its size is unstable. <br/>
time.Timer       80 bytes   // it is a struct, so its size is unstable. <br/>
time.Duration    8 bytes <br/>
[]byte           24 bytes <br/>


## The above code run result is:
Size of main.Employee1 struct: 216 bytes <br/>
Size of main.Employee2 struct: 192 bytes <br/>
Size of main.Employee3 struct: 208 bytes <br/>
