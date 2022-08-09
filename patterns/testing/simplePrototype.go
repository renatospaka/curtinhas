package testing

import (
	"fmt"

	"github.com/renatospaka/design-patterns/creational/prototype"
)

func TestingPrototype() {
	fmt.Println("This is the PROTOTYPE implementation")

	file1 := &prototype.File{Name: "File1"}
	file2 := &prototype.File{Name: "File2"}
	file3 := &prototype.File{Name: "File3"}

	folder1 := &prototype.Folder{
		Children: []prototype.NodeInterface{file1},
		Name:     "Folder1",
	}

	folder2 := &prototype.Folder{
		Children: []prototype.NodeInterface{folder1, file2, file3},
		Name:     "Folder2",
	}
	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.Print("  ")

	cloneFolder := folder2.Clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.Print("  ")
	
	fmt.Println("==================================")
	fmt.Println()
}
