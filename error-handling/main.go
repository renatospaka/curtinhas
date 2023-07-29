package main

import (
	"errors"
	"fmt"
	"io"
	"reflect"
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
		fmt.Printf("Has ErrOne of type %v \n", reflect.TypeOf(ErrOne))
		fmt.Printf("Has ErrOne of type %v \n", reflect.ValueOf(ErrOne).Kind())
	}
	fmt.Println()

	if errors.Is(anotherJoin, ErrTwo) {
		fmt.Printf("Has ErrTwo of type %T \n", ErrTwo)
		fmt.Printf("Has ErrTwo of type %v \n", reflect.TypeOf(ErrTwo))
		fmt.Printf("Has ErrTwo of type %v \n", reflect.ValueOf(ErrTwo).Kind())
	}
	fmt.Println()

	if errors.Is(anotherJoin, errThree) {
		fmt.Printf("Has errThree of type %T \n", errThree)
		fmt.Printf("Has errThree of type %v \n", reflect.TypeOf(errThree))
		fmt.Printf("Has errThree of type %v \n", reflect.ValueOf(errThree).Kind())
	}
	fmt.Println()

	if errors.Is(anotherJoin, errFour) {
		fmt.Printf("Has errFour of type %T \n", errFour)
		fmt.Printf("Has errFour of type %v \n", reflect.TypeOf(errFour))
		fmt.Printf("Has errFour of type %v \n", reflect.ValueOf(errFour).Kind())
	}
	fmt.Println()

	if errors.Is(anotherJoin, errFive) {
		fmt.Printf("Has errFive of type %T \n", errFive)
		fmt.Printf("Has errFive of type %v \n", reflect.TypeOf(errFive))
		fmt.Printf("Has errFive of type %v \n", reflect.ValueOf(errFive).Kind())
	}
	fmt.Println()

	if errors.Is(anotherJoin, ErrSix) {
		fmt.Printf("Has ErrSix of type %T \n", ErrSix)
		fmt.Printf("Has ErrSix of type %v \n", reflect.TypeOf(ErrSix))
		fmt.Printf("Has ErrSix of type %v \n", reflect.ValueOf(ErrSix).Kind())
	}
	fmt.Println()

	// Using fmt.Errorf with multiple %w
	efmt := fmt.Errorf("--> %w %w", errors.New("fmt"), io.EOF)	
	fmt.Printf("==> %+v\n", efmt)

	if errors.Is(efmt, io.EOF) {
		fmt.Println("End of File")
	}
}
