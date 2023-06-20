package sample 

import (
	"testing"
)
const (
	SEP = "\n__________________________________________________________________________________\n"
)
func Test_1 (* testing.T) () {

	run()
}

func Test_2 (* testing.T) () {
	log(SEP)
	pos := 6
	len := 13 
	v1 := 100 
	v2 , err := B(pos , len , v1)
	log(SEP)
	log("v2 : %v , err : %v " , v2 , err)
}

func Test_3 (* testing.T) () {
	log(SEP)
	pos := 0
	len := 1 
	v1 := 0xFF 
	v2 , err := B(pos , len , v1)
	log(SEP)
	log("v2 : %v , err : %v " , v2 , err)
}