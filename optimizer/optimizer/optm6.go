package optimizer

import (
	"fmt"
	"log"
	"time"
)

type str struct {
	word string
}

func Main6() {
	fmt.Println("#6 - Don’t use an empty interface in a map!")
	defineMapFirstWay()
	defineMapSecWay()
	defineMapThirdWay()
	defineMapFourthWay()
}

func defineMapFirstWay() {
	execTime := time.Now()
	words := make(map[int]interface{})
	words[1] = "7"
	words[2] = "simple"
	words[3] = "ways"
	words[4] = "to"
	words[5] = "optimise"
	words[6] = "Golang"
	_ = words
	log.Printf("defineMapFirstWay ExecTime is %21s %6s\n", "->", time.Since(execTime))
}

func defineMapSecWay() {
	execTime := time.Now()
	var words = map[int]str{}
	words[1] = str{"7"}
	words[2] = str{"simple"}
	words[3] = str{"ways"}
	words[4] = str{"to"}
	words[5] = str{"optimise"}
	words[6] = str{"Golang"}
	_ = words
	log.Printf("defineMapSecWay ExecTime is %23s %6s\n", "->", time.Since(execTime))
}

func defineMapThirdWay() {
	execTime := time.Now()
	words := make(map[int]string)
	words[1] = "7"
	words[2] = "simple"
	words[3] = "ways"
	words[4] = "to"
	words[5] = "optimise"
	words[6] = "Golang"
	_ = words

	log.Printf("defineMapThirdWay ExecTime is %21s %6s\n", "->", time.Since(execTime))
}

func defineMapFourthWay() {
	execTime := time.Now()
	words := map[int]string{1: "7", 2: "simple", 3: "ways", 4: "to", 5: "optimise", 6: "Golang"}

	_ = words
	log.Printf("defineMapFourthWay ExecTime is %20s %6s\n", "->", time.Since(execTime))
}