package main

import "github.com/renatospaka/design-patterns/testing"

func main() {
	testingCreationalPatterns()
	testingStructuralPatterns()
}

func testingCreationalPatterns() {
	testing.TestingAbstractFactory()
	testing.TestingBuilder()
	testing.TestingFactory()
}

func testingStructuralPatterns() {
	testing.TestingFacade()
}
