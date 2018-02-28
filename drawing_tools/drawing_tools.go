package drawing_tools 

import (
		"fmt"
		"encoding/json"
		"io/ioutil"
		"os"
		"sort"
       )


type Parameters struct {
	label			string		`json:"label"`
}

type Operations []Operation

func (o Operations) Len() int {
	return len(o)
}

func (o Operations) Less(i, j int) bool{
	return Depth(o[i],0) > Depth(o[j],0)
}

func (o Operations) Swap(i, j int){
	o[i],o[j] = o[j], o[i]
}

func Depth(o Operation, currentDepth int) int {
	d := currentDepth 
	if o.Children!=nil {
		for _,child := range o.Children {
			dt := Depth(child,d+1)
			if d<dt {

				d = d + dt

			}
		}
	}else{
		return d
	}
	return d
}


type Operation struct {
	Label			string		`json:"label"`
	Type_op			string		`json:"type_op"`
	Condition_False		*Operation	`json:"condition_false,omitempty"`
	Condition_True		*Operation	`json:"condition_true,omitempty"`
	Depends_On		*Operation	`json:"depends_on,omitempty"`
	Children		Operations	`json:"children"`
//	node			string		`json:"node"`
	target			string		`json:"target,omitempty"`
//	expectedDuration	string		`json:"expectedDuration"`
//	params			string		`json:"params"`
}

// Parsing function to extract JSON
// Got from https://www.chazzuka.com/2015/03/load-parse-json-file-golang/
func Parse(filename string) Operation {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c Operation
	json.Unmarshal(raw, &c)
	return c
}


func PrintArray(array [][]string){
	
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

	// expanding output array
	for i:=0; i<len(box); i++{
		if len(output[i+level])<step+len(box[i]){
			lengthOutput := len(output[i+level])
			for j:=0; j<(step+len(box[i])-lengthOutput+1); j++{

				output[i+level] = append(output[i+level]," ")
			}
		}
	}

	// copying box in output
	for i:=0; i<len(box);i++{
		for j:= 0; j< len(box[i]);j++{
			output[level + i][step + j] = box[i][j]
		}
	}

	return level,step+len(box[0]),output
}

// This function draws a blank to fill the gap in a fork.
//
// INPUTS
// level : vertical strip (top left corner)
// step : final length needed
// output: array of characters of the box before drawing
//
// OUTPUTS
// tree with the new box
func addBlank(level int, stepBefore int, stepFinal int, output [][]string) ([][]string){

	// arbitrary length of the box
	// TODO change with the actual size.
	
	length := stepFinal-stepBefore
	
	height := 4
	box :=  make([][]string,height)

	if length != 0 {



		// init empty box

		for i:= 0; i<height; i++ {
			// 8 = 5 for the arrow, 3 for the final O
			box[i] = make([]string,length)
			for j:=0; j<len(box[i]);j++{
				box[i][j]=" "
			}

		}

		_,_,output = drawBox(level,stepBefore,box,output)
	}


	return output
}

// This function draws a IF box.
//
// INPUTS
// level : vertical strip (top left corner)
// levelFinal: final length needed
// step : vertical blablabla
// output: array of characters of the box before drawing
//
// OUTPUTS
// tree with the new box
func addConditionalBox(level int, levelFinal int, step int, output [][]string) (int,int,[][]string){

	// arbitrary length of the box
	// TODO change with the actual size.
	
	height := levelFinal-level+4
	
	length := 15
	box :=  make([][]string,height)


		// init empty box

	for i:= 0; i<height; i++ {
		// 8 = 5 for the arrow, 3 for the final O
		box[i] = make([]string,length)
		for j:=0; j<len(box[i]);j++{
			box[i][j]=" "
		}

	}

	for j:=0; j<len(box[1]);j++{
		box[1][j]="─"
	}
	box[1][4]="┬"
	box[1][5]="["
	box[1][6]=" "
	box[1][7]="S"
	box[1][8]="I"
	box[1][9]=" "
	box[1][10]="]"


	for j:=4; j<len(box[1]);j++{
		box[levelFinal+1][j]="─"
	}

	box[height-3][4]="└"
	box[height-3][5]="["
	box[height-3][6]="S"
	box[height-3][7]="I"
	box[height-3][8]="N"
	box[height-3][9]="O"
	box[height-3][10]="N"
	box[height-3][11]="]"

	fmt.Printf("height %d\n",height)
	fmt.Printf("box %d\n",len(box))
	i:=height-4
	for box[i][3]!="─"{
		box[i][4]="│"
		i--
	}

	level,step,output = drawBox(level,step,box,output)

	return level,step,output
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



	if length != 0 {



		// init empty box

		for i:= 0; i<height; i++ {
			// 8 = 5 for the arrow, 3 for the final O
			box[i] = make([]string,length)
			for j:=0; j<len(box[i]);j++{
				box[i][j]="."
			}

		}

		for j:=0; j<len(box[1])-1;j++{
			box[1][j]="─"
		}

		//box[1][len(box[1])-1]=">"

		_,_,output = drawBox(level,stepBefore,box,output)
	}


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

	box :=  make([][]string,height)
	for i:= 0; i<height; i++ {
		// 8 = 5 for the arrow, 3 for the final O
		box[i] = make([]string,length+6)
		for j:=0; j<len(box[i]);j++{
			box[i][j]=" "
		}

	}

	box[0][4] = "┌"
	// drawing the top
	for i:=5; i<length+3+1; i++{
		box[0][i] = "─"
	}
	box[0][length+3+1] = "┐"
	
	
	// drawing bottom
	box[3][4] = "└"
       	for i:=5; i<length+3+1; i++{
		box[3][i]="─"
	}
	box[3][length+3+1] = "┘"
	
	// drawing left 
	box[2][4] = "│"
	box[1][4] = "│"

	// drawing right
	box[1][4+length] = "│"
	box[2][4+length] = "│"

	// drawing 'in' arrow
	box[1][4-1] = "→"
	box[1][4-2] = "─"
	box[1][4-3] = "─"
	box[1][4-4] = "─"

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
func DrawGraph(level int, step int, operations []Operation, output [][]string) (int, int, [][]string) { 

	for _,operation := range operations {

		if operation.Type_op == "operation" {
			level,step,output = addBox(level,step,operation,output)
		}

		if operation.Type_op == "sequence"  {	
		
			if operation.Children != nil {
				level,step,output = DrawGraph(level,step,operation.Children,output)
			}
		}

		if operation.Type_op == "fork" {
			
			stepBefore := step

			maxStep := step
			tmp_step :=0
			tmp_level := level-4
			sort.Sort(Operations(operation.Children))

			for _,child := range operation.Children {
				 
				tmp_level,tmp_step,output = DrawGraph(tmp_level+4,step+1,[]Operation{child},output)

				if tmp_step > maxStep {
					maxStep=tmp_step
				}

				// Filling lines
				addLine(tmp_level,tmp_step,maxStep+1,output)
				output[tmp_level][maxStep]="!"
				output[tmp_level+1][maxStep]="!"
				output[tmp_level+1][maxStep-1]=">"
				output[tmp_level+2][maxStep]="!"
				output[tmp_level+3][maxStep]="!"

			}
		
			i := tmp_level +1
			// Opening bracket

			output[level][stepBefore]="F"	

			for output[i-1][stepBefore] != "F" {
				output[i][stepBefore]="│"
				i--
			}
			output[i][stepBefore]="┬"
			
			step = maxStep +1
			level = tmp_level
		}

		if operation.Type_op == "condition" {
			// TODO conditionnal
			tmp_level := 0
			tmp_step := 0


			tmp_level,tmp_step,output = DrawGraph(level,step+14,[]Operation{*operation.Condition_True},output)
			//output[level+1][step-3]="S"
			//output[level+1][step-2]="I"
			tmp_level,tmp_step,output = DrawGraph(tmp_level+4,step+14,[]Operation{*operation.Condition_False},output)

			tmp_level,tmp_step,output = addConditionalBox(level,tmp_level,step,output)


			step = tmp_step
			level = tmp_level
		}

	}
	return level, step, output

}


