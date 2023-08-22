package main

import (
	"errors"
	"fmt"
	"io"
)


var ErrOne = fmt.Errorf("One")
var ErrTwo = fmt.Errorf("Two")
var ErrSix = fmt.Errorf("Six")

func main() {
	join := errors.Join(ErrOne, ErrTwo)

	errThree := fmt.Errorf("Three")
	fmt.Printf("%v\n", join)
	fmt.Println()

	newJoin := errors.Join(join, errThree)
	fmt.Printf("%v\n", newJoin)
	fmt.Println()

	errFour := fmt.Errorf("Four")
	errFive := fmt.Errorf("Five")
	anotherJoin := errors.Join(newJoin, errFour, errFive)
	fmt.Printf("%v\n", anotherJoin)
	fmt.Println()

	if errors.Is(anotherJoin, ErrOne) {
		fmt.Printf("Has ErrOne of type %T \n", ErrOne)
		// fmt.Printf("Has ErrOne of type %v \n", reflect.TypeOf(ErrOne))
		// fmt.Printf("Has ErrOne of type %v \n", reflect.ValueOf(ErrOne).Kind())
	}
	fmt.Println()

	if errors.Is(anotherJoin, ErrTwo) {
		fmt.Printf("Has ErrTwo of type %T \n", ErrTwo)
		// fmt.Printf("Has ErrTwo of type %v \n", reflect.TypeOf(ErrTwo))
		// fmt.Printf("Has ErrTwo of type %v \n", reflect.ValueOf(ErrTwo).Kind())
	}
	fmt.Println()

	if errors.Is(anotherJoin, errThree) {
		fmt.Printf("Has errThree of type %T \n", errThree)
		// fmt.Printf("Has errThree of type %v \n", reflect.TypeOf(errThree))
		// fmt.Printf("Has errThree of type %v \n", reflect.ValueOf(errThree).Kind())
	}
	fmt.Println()

	if errors.Is(anotherJoin, errFour) {
		fmt.Printf("Has errFour of type %T \n", errFour)
		// fmt.Printf("Has errFour of type %v \n", reflect.TypeOf(errFour))
		// fmt.Printf("Has errFour of type %v \n", reflect.ValueOf(errFour).Kind())
	}
	fmt.Println()

	if errors.Is(anotherJoin, errFive) {
		fmt.Printf("Has errFive of type %T \n", errFive)
		// fmt.Printf("Has errFive of type %v \n", reflect.TypeOf(errFive))
		// fmt.Printf("Has errFive of type %v \n", reflect.ValueOf(errFive).Kind())
	}
	fmt.Println()

	if errors.Is(anotherJoin, ErrSix) {
		fmt.Printf("Has ErrSix of type %T \n", ErrSix)
		// fmt.Printf("Has ErrSix of type %v \n", reflect.TypeOf(ErrSix))
		// fmt.Printf("Has ErrSix of type %v \n", reflect.ValueOf(ErrSix).Kind())
	}
	fmt.Println()

	// Using fmt.Errorf with multiple %w
	efmt := fmt.Errorf("--> %w %w", errors.New("fmt"), io.EOF)	
	fmt.Printf("==> %+v\n", efmt)

	if errors.Is(efmt, io.EOF) {
		fmt.Println("End of File")
	}
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()


	var e error = nil
	// e0 := errors.New("this is Err0")
	e = errors.Join(e, errors.New("this is Err0"))
	// e1 := errors.New("this is Err1")
	e = errors.Join(e, errors.New("this is Err1"))
	// e2 := errors.New("this is Err2")
	e = errors.Join(e, errors.New("this is Err2"))
	// e3 := errors.New("this is Err3")
	e = errors.Join(e, errors.New("this is Err3"))
	fmt.Println("start reading all errors in e")
	fmt.Printf("first let's check how many errors there are there: \n%v\n", e)
	fmt.Println("reading...")
	var errs []error = []error{}
	errs = UnwrapJoin(e)
	for x := 0; x < len(errs); x++ {
		fmt.Printf("%dth error = %v\n", x+1, errs[x])
	}
	fmt.Println("read!")

	// wrapped := errors.Unwrap(e)
	fmt.Printf("Unrapped errors: %v\n", errors.Unwrap(e))
}

type Unwrapper interface {
	Unwrap() []error
}

func UnwrapJoin(err error) []error {
	if err == nil {
		return nil
	}
	errs := []error{}
	if err, ok := err.(Unwrapper); ok {
		for _, e := range err.Unwrap() {
			errs = append(errs, UnwrapJoin(e)...)
		}
		return errs
	}
	return append(errs, err)

}
