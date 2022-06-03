package optimizer

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func Main5() {
	fmt.Println("#5 - Instead for range with for loop!")
	s := [11]string{"6", " simple", " ways", " to", " optimise", " Golang", " -", " this", " is", " the", " catch"}

	str0 := joinStrWithSprintf(s)
	log.Println(str0)

	str1 := joinStrWithBuilderForRange(s)
	log.Println(str1)

	str2 := joinStrWithBuilderForLoop(s)
	log.Println(str2)
}

func joinStrWithSprintf2(str [11]string) (allStr string) {
	// just to keep things similar
	joinStrWithSprintf(str)
	return
}

func joinStrWithBuilderForRange(str [11]string) (allStr string) {
	execTime := time.Now()

	var sb strings.Builder
	for _, s := range str {
		sb.WriteString(s)
	}

	log.Printf("joinStrWithBuilder ExecTime is %29s %6s\n", "->", time.Since(execTime))
	return sb.String()
}

func joinStrWithBuilderForLoop(s [11]string) (allStr string) {
	execTime := time.Now()

	var sb strings.Builder
	for i := 0; i < len(s); i++ {
		sb.WriteString(s[i])
	}

	log.Printf("joinStrWithBuilderForLoop ExecTime is %22s %6s\n", "->", time.Since(execTime))
	return sb.String()
}