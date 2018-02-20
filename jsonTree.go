package main 

import (
		"fmt"
		. "github.com/clemshubs/jsonTree/drawing_tools"
       )

func main() {

	
	// TODO : change this fixed array with dynamic allocation (append,...)
	output :=  make([][]string,4)

	fmt.Printf("\n")
	fmt.Printf("Loading file and parsing\n")

	// Parsing of the JSON
	var jsonContent Operation

	jsonContent = Parse("D:/Utilisateurs/hubinac/Documents/testCondGOO.goo")

	// Computation of the tree
	_,_,output = DrawGraph(0,0,[]Operation{jsonContent},output)

	// Drawing of the tree
	PrintArray(output)

	// Bye
	fmt.Printf("End of program")
}
