package main 
import (
		"fmt"
		"testing"
		. "github.com/clemshubs/jsonTree/drawing_tools"
       )

func TestDrawOneBox(t  *testing.T){
	if testing.Short(){
		t.Skip("skippig test in shor mode.")
	}

	// init output array
	output :=  make([][]string,1)

	// init box array
	box :=  make([][]string,4)

	box[0] = append(box[0]," "," "," "," ","-","-","-","-","-","-","-","-","-","-")
	box[1] = append(box[1],"-","-","-",">","|","toto","","",""," "," "," "," ","|")
	box[2] = append(box[2]," "," "," "," ","|","titi","","",""," "," "," "," ","|")
	box[3] = append(box[3]," "," "," "," ","-","-","-","-","-","-","-","-","-","-")

	_,_,output = DrawBox(0,0,box,output)

	for i := 0; i < len(box); i++ {
		for j := 0; j < len(box[i]);j++{
			if box[i][j] != output[i][j]{
				fmt.Print("Drawing expected\n")
				PrintArray(box)
				fmt.Print("Drawing found")
				PrintArray(output)
				t.Error("One box drawing failed")
			}
		}
	}



}

func TestDrawTwoBoxes(t  *testing.T){
	if testing.Short(){
		t.Skip("skippig test in shor mode.")
	}

	// init output array
	output :=  make([][]string,4)

	// init box array
	box :=  make([][]string,4)

	box[0] = append(box[0]," "," "," "," ","-","-","-","-","-","-","-","-","-","-")
	box[1] = append(box[1],"-","-","-",">","|","toto","","",""," "," "," "," ","|")
	box[2] = append(box[2]," "," "," "," ","|","titi","","",""," "," "," "," ","|")
	box[3] = append(box[3]," "," "," "," ","-","-","-","-","-","-","-","-","-","-")

	step := 0
	// init box array
	box1 :=  make([][]string,4)

	box1[0] = append(box1[0]," "," "," "," ","-","-","-","-","-","-","-","-","-","-")
	box1[1] = append(box1[1],"-","-","-",">","|","tata","","",""," "," "," "," ","|")
	box1[2] = append(box1[2]," "," "," "," ","|","titi","","",""," "," "," "," ","|")
	box1[3] = append(box1[3]," "," "," "," ","-","-","-","-","-","-","-","-","-","-")

	goodOutput := make([][]string,4)
	goodOutput[0] = append(goodOutput[0],box[0]...)
	goodOutput[0] = append(goodOutput[0],box1[0]...)
	goodOutput[1] = append(goodOutput[1],box[1]...)
	goodOutput[1] = append(goodOutput[1],box1[1]...)
	goodOutput[2] = append(goodOutput[2],box[2]...)
	goodOutput[2] = append(goodOutput[2],box1[2]...)
	goodOutput[3] = append(goodOutput[3],box[3]...)
	goodOutput[3] = append(goodOutput[3],box1[3]...)

	PrintArray(goodOutput)

	_,step,output = DrawBox(0,step,box,output)
	_,step,output = DrawBox(0,step,box,output)

	for i := 0; i < len(box); i++ {
		for j := 0; j < len(box[i]);j++{
			if box[i][j] != output[i][j]{
				fmt.Print("Drawing expected\n")
				PrintArray(goodOutput)
				fmt.Print("Drawing found")
				PrintArray(output)
			t.Error("One box drawing failed")
			}
		}
	}



}
