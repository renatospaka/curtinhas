package optimizer

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func Main4() {
	fmt.Println("#4 - Join string using strings.Builder!")
	s := [11]string{"6", " simple", " ways", " to", " optimise", " Golang", " -", " this", " is", " the", " catch"}
	str0 := joinStrDirect(s)
	log.Println(str0)

	str1 := joinStrWithSprintf(s)
	log.Println(str1)

	str2 := joinStrWithBuilder(s)
	log.Println(str2)
}

func joinStrDirect(str [11]string) (allStr string) {
	execTime := time.Now()

	for _, s := range str {
		allStr += s
	}

	log.Printf("joinStrDirect ExecTime is %22s %6s\n", "->", time.Since(execTime))
	return
}

func joinStrWithSprintf(str [11]string) (allStr string) {
	execTime := time.Now()

	for _, s := range str {
		allStr += fmt.Sprintf("%s", s)
	}

	log.Printf("joinStrWithSprintf ExecTime is %17s %6s\n", "->", time.Since(execTime))
	return
}

func joinStrWithBuilder(str [11]string) (allStr string) {
	execTime := time.Now()

	var sb strings.Builder
	for _, s := range str {
		sb.WriteString(s)
	}

	log.Printf("joinStrWithBuilder ExecTime is %17s %6s\n", "->", time.Since(execTime))
	return sb.String()
}
