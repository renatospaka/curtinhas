package optimizer

import (
	"fmt"
	"time"
	"unsafe"
)

/**
In golang, data type and size list is following:
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
*/

type commonStruct struct {
	createAt time.Time
	updateAt time.Timer
	timeout  time.Duration
	jsonStr  []byte
}

type employee1 struct {
	IsActive  bool
	Age       int64
	IsMarried bool
	Name      string
	weight    int32
	height    int16
	PhotoLen  float64
	PhotoWid  float32
	intNum    int
	length    int8
	common    commonStruct
}

type employee2 struct {
	Name      string
	Age       int64
	intNum    int
	PhotoLen  float64
	PhotoWid  float32
	weight    int32
	height    int16
	length    int8
	IsMarried bool
	IsActive  bool
	common    commonStruct
}

type employee3 struct {
	Name      string
	IsActive  bool
	intNum    int
	length    int8
	Age       int64
	IsMarried bool
	weight    int32
	height    int16
	PhotoWid  float32
	PhotoLen  float64
	common    commonStruct
}

var empl1 employee1
var empl2 employee2
var empl3 employee3

func Main() {
	fmt.Printf("Size of %T struct: %d bytes\n", empl1, unsafe.Sizeof(empl1))

	fmt.Printf("Size of %T struct: %d bytes\n", empl2, unsafe.Sizeof(empl2))

	fmt.Printf("Size of %T struct: %d bytes\n", empl3, unsafe.Sizeof(empl3))
}
