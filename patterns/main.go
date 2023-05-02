package main

import "github.com/renatospaka/design-patterns/testing"

func main() {
	testingCreationalPatterns()
	testingStructuralPatterns()
	testingBehavioralPatterns()
}

func testingCreationalPatterns() {
	testing.TestingAbstractFactory()
	testing.TestingBuilder()
	testing.TestingFactory()
	testing.TestingPrototype()
}

func testingStructuralPatterns() {
	testing.TestingFacade()
}

func testingBehavioralPatterns() {
	testing.ChainOfResponsability()
}
