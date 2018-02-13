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

func drawBox(level int, step int, operation Operation, output [][]string) (int, int, [][]string){
	length := 10
	fmt.Printf(operation.Label)	
	fmt.Printf("\n")
	fmt.Printf("level=%d\n",level)	
	fmt.Printf("step=%d\n",step)	
	fmt.Printf("\n")

	for i:=step; i<step+length+3; i++{
		output[level][i] = "-"
	}

	output[level+1][step+1]=operation.Type_op
	output[level+2][step+1]=operation.Label

	for i:= step+2; i< step+1+len(operation.Type_op); i++{
		output[level+1][i]=""
	}
	for i:= step+2; i< step+1+len(operation.Label); i++{
		output[level+2][i]=""
	}
	
	output[level+1][step] = "|"
	output[level+1][step+length+2] = "|"
	output[level+1][step-1] = ">"
	output[level+1][step-2] = "-"
	output[level+1][step-3] = "-"

	output[level+2][step] = "|"
	output[level+2][step+length+2] = "|"
	for i:=step; i<step+length+3; i++{
		output[level+3][i]="-"
	}
	
	step = step + length + 7

	return level,step,output

}

// Print a list of operations for a level and a step
func drawGraph(level int, step int, operations []Operation, output [][]string) (int, int, [][]string) { 

	for _,operation := range operations {
	
		if operation.Type_op == "operation" {
			level,step,output = drawBox(level,step,operation,output)
		}

		if operation.Type_op == "sequence"  {	
		
			if operation.Children != nil {
				level,step,output = drawGraph(level,step,operation.Children,output)
			}
		}

		if operation.Type_op == "fork" {
			child := operation.Children [0]
			
			stepBefore := step
			level,_,output = drawGraph(level,step,[]Operation{child},output)


			 for _,child := range operation.Children[1:] {
			 	 level=level+4
				 
				 output[level-2][stepBefore-3] = "|"	
				 output[level-1][stepBefore-3] = "|"	
				 output[level][stepBefore-3] = "|"	
				 output[level+1][stepBefore-3] = "|"	
				 output[level+1][stepBefore-2] = "-"	
				 output[level+1][stepBefore-1] = ">"
				 level,step,output = drawGraph(level,stepBefore,[]Operation{child},output)
			 }
		}

		if operation.Type_op == "conditionnal" {
		
		
		}

	}
	return level, step, output

}


func main() {

	

	output :=  make([][]string,50)
	for i:= 0; i<50; i++ {
		output[i] = make([]string,100)
	}
	for i := 0; i < len(output); i++ {
		for j := 0; j < len(output[i]);j++{
			output[i][j]=" "
		}
	}
	output[1][0]="-"
	output[1][1]="-"
	output[1][2]="-"
	output[1][3]="-"
	output[1][4]=">"

	fmt.Printf("\n")
	fmt.Printf("Loading file and parsing\n")
	var jsonContent Operation

	jsonContent = parse("D:/Utilisateurs/hubinac/Documents/test.goo")
	_,_,output = drawGraph(0,5,[]Operation{jsonContent},output)
	for i := 0; i < len(output); i++ {
		for j := 0; j < len(output[i]);j++{
			fmt.Print(output[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("End of program")
}
