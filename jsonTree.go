package main

import (
		"fmt"
		"encoding/json"
		"io/ioutil"
		"os"
       )

type Parameters struct {
	label			string		`json:"label"`
}

type Operation struct {
	Label			string		`json:"label"`
	Type_op			string		`json:"type_op"`
	Condition_False		*Operation	`json:"condition_false,omitempty"`
	Condition_True		*Operation	`json:"condition_true,omitempty"`
	Children		[]Operation	`json:"children"`
//	node			string		`json:"node"`
//	target			string		`json:"target"`
//	expectedDuration	string		`json:"expectedDuration"`
//	params			string		`json:"params"`
}

// Parsing function to extract JSON
// Got from https://www.chazzuka.com/2015/03/load-parse-json-file-golang/
func parse(filename string) Operation {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c Operation
	json.Unmarshal(raw, &c)
	return c
}

func printArray(array [][]string){
	
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]);j++{
			fmt.Print(array[i][j])
		}
		fmt.Printf("\n")
	}

}

// This function returns the string array with the new box inserted
// 
// INPUTS
// level : vertical position of top left corner
// step : horizontal position of top left corner
// box : string array of the box
// output : the complete array of strings
//
// OUPUTS
// the array with the new box
func drawBox(level int, step int, box [][]string, output [][]string)(int,int,[][]string){

	// if a line doesn't exist, we add it
	if len(output)<level+4 {
		outputL := len(output)
		for i:=0; i<level+4-outputL; i++ {
			output = append(output,[]string{})
		}
	}

	fmt.Print("len output ")
	fmt.Print(len(output))
	fmt.Print("\n")
	fmt.Print("level ")
	fmt.Print(level )
	fmt.Print("\n")
	fmt.Print("step ")
	fmt.Print(step)
	fmt.Print("\n")
	
	// expanding output array
	for i:=0; i<len(box); i++{
		if len(output[i+level])<step+len(box[i]){
			lengthOutput := len(output[i+level])
			for j:=0; j<(step+len(box[i])-lengthOutput+1); j++{

				output[i+level] = append(output[i+level]," ")
			}
		}
	}

	fmt.Print("len output ")
	fmt.Print(len(output[level]))
	fmt.Print("\n")
	fmt.Print("level ")
	fmt.Print(level)
	fmt.Print("\n")


	// copying box in output
	for i:=0; i<len(box);i++{
		for j:= 0; j< len(box[i]);j++{
			output[level + i][step + j] = box[i][j]
		}
	}


	return level,step+len(box[0]),output
}


// This function draws a line to fill the gap in a fork.
//
// INPUTS
// level : vertical strip (top left corner)
// step : final length needed
// output: array of characters of the box before drawing
//
// OUTPUTS
// tree with the new box
func addLine(level int, stepBefore int, stepFinal int, output [][]string) ([][]string){

	// arbitrary length of the box
	// TODO change with the actual size.
	
	length := stepFinal-stepBefore
	
	height := 4
	box :=  make([][]string,height)
	fmt.Printf("level=%d\n",level)	
	fmt.Printf("stepBefore=%d\n",stepBefore)	
	fmt.Printf("stepFinal=%d\n",stepFinal)	
	fmt.Printf("\n")



	if length != 0 {



		// init empty box

		for i:= 0; i<height; i++ {
			// 8 = 5 for the arrow, 3 for the final O
			box[i] = make([]string,length)
			for j:=0; j<len(box[i]);j++{
				box[i][j]=" "
			}

		}

		for j:=0; j<len(box[1])-1;j++{
			box[1][j]="-"
		}

		//box[1][len(box[1])-1]=">"

		_,_,output = drawBox(level,stepBefore,box,output)
	}


	printArray(output)
	return output
}


// This function draws a box for an operation from its top left corner location
//
// INPUTS
// level : vertical strip (top left corner)
// step : horizontal strip (top left corner)
// operation : operation to draw
// output: array of characters of the box before drawing
//
// OUTPUTS
// top left corner vertical strip
// bottom right corner horizontal strip
// tree with the new box
func addBox(level int, step int, operation Operation, output [][]string) (int, int, [][]string){

	// arbitrary length of the box
	// TODO change with the actual size.
	
	length := 10
	height := 4

	fmt.Printf(operation.Label)	
	fmt.Printf("\n")
	fmt.Printf("level=%d\n",level)	
	fmt.Printf("step=%d\n",step)	
	fmt.Printf("\n")

	// TODO : change this fixed array with dynamic allocation (append,...)
	box :=  make([][]string,height)
	for i:= 0; i<height; i++ {
		// 8 = 5 for the arrow, 3 for the final O
		box[i] = make([]string,length+6)
		for j:=0; j<len(box[i]);j++{
			box[i][j]=" "
		}

	}

	// drawing the top
	for i:=4; i<length+3+2; i++{
		box[0][i] = "-"
	}

	// drawing bottom
       	for i:=4; i<length+3+2; i++{
		box[3][i]="-"
	}
	
	// drawing left 
	box[2][4] = "|"
	box[1][4] = "|"

	// drawing right
	box[1][4+length] = "|"
	box[2][4+length] = "|"

	// drawing 'in' arrow
	box[1][4-1] = ">"
	box[1][4-2] = "-"
	box[1][4-3] = "-"
	box[1][4-4] = "-"

	// dranwing 'final' o (to maybe be overwriten by next arrow)
	box[1][4+length+1] = "o"

	// writing info
	box[1][4+1]=operation.Type_op
	box[2][4+1]=operation.Label

	for i:= 4+2; i< 4+1+len(operation.Type_op); i++{
		box[1][i]=""
	}
	for i:= 4+2; i< 4+1+len(operation.Label); i++{
		box[2][i]=""
	}
	
	return drawBox(level,step,box,output)
}

// Functional drawing of the graph
// 
// INPUTS
//
// OUTPUTS
//
func drawGraph(level int, step int, operations []Operation, output [][]string) (int, int, [][]string) { 

	for _,operation := range operations {


//		fmt.Print(operation.Label)
		fmt.Print("\n")
		if operation.Type_op == "operation" {
			level,step,output = addBox(level,step,operation,output)
		}

		if operation.Type_op == "sequence"  {	
		
			if operation.Children != nil {
				level,step,output = drawGraph(level,step,operation.Children,output)
			}
		}

		if operation.Type_op == "fork" {
			child := operation.Children [0]
			

			
			stepBefore := step
			levelBefore := level
			maxStep := step
			tmp_step :=0
			
			level,tmp_step,output = drawGraph(level,step,[]Operation{child},output)

			if tmp_step > maxStep {
				maxStep=tmp_step
			}

			for _,child := range operation.Children[1:] {
			 	 level=level+4
				 
				 level,tmp_step,output = drawGraph(level,stepBefore,[]Operation{child},output)
				 output[level-3][stepBefore] = "|"	
				 output[level-2][stepBefore] = "|"	
				 output[level-1][stepBefore] = "|"	
				 output[level][stepBefore] = "|"	
				 output[level+1][stepBefore] = "|"
				if tmp_step > maxStep {
					maxStep=tmp_step
				}
			}
		
			fmt.Printf("LEVEL %d",levelBefore)
			fmt.Printf("STEP BEFORE %d\n",stepBefore)
			i := level

			// Opening bracket

			output[levelBefore][stepBefore]="F"	
			for output[i][stepBefore] != "F" {
				output[i][stepBefore]="|"
				output[i-1][stepBefore]="|"
				output[i-2][stepBefore]="|"
				output[i-3][stepBefore]="|"
				i=i-4
			}

			// Filling lines
			for i:=levelBefore; i<level; i=i+4{
				addLine(i,len(output[i])-2,maxStep+1,output)
			}

			// Closing bracket
			for i:=levelBefore+2;i<level+1;i++ {
				output[i][maxStep-1]="^"
			}
			printArray(output)
			level = levelBefore
			step = maxStep

		}

		if operation.Type_op == "conditionnal" {
		
		
		}

	}
	return level, step, output

}


func main() {

	
	// TODO : change this fixed array with dynamic allocation (append,...)
	output :=  make([][]string,4)

	fmt.Printf("\n")
	fmt.Printf("Loading file and parsing\n")

	// Parsing of the JSON
	var jsonContent Operation

	jsonContent = parse("D:/Utilisateurs/hubinac/Documents/test.goo")

	// Computation of the tree
	_,_,output = drawGraph(0,0,[]Operation{jsonContent},output)

	// Drawing of the tree
	printArray(output)

	// Bye
	fmt.Printf("End of program")
}
