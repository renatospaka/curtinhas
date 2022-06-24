package main

import "github.com/renatospaka/design-patterns/testing"

func main() {
	testingCreationalPatterns()
	testingStructuralPatterns()
}

func testingCreationalPatterns() {
	testing.TestingAbstractFactory()
}

func testingStructuralPatterns() {
	testing.TestingFacade()
}
