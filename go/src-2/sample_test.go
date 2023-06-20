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

func Test_2 (t * testing.T) () {
	log(SEP)
	pos := 6
	len := 13 
	v1 := 100 
	v2 , err := B(pos , len , v1)
	log(SEP)
	log("v2 : %v , err : %v  , pos:%v , len:%v , v1:%v\n" , v2 , err , pos , len , v1)
	if v2 != 100 {
		t.Errorf("Failed.")
	}
}

func Test_3 (t * testing.T) () {
	log(SEP)
	pos := 0
	len := 1 
	v1 := 0xFF 
	v2 , err := B(pos , len , v1)
	log(SEP)
	log("v2 : %v , err : %v  , pos:%v , len:%v , v1:%v\n" , v2 , err , pos , len , v1)
	if v2 != 1 {
		t.Errorf("Failed.")
	}
}

func Test_4 (t * testing.T) () {
	log(SEP)
	pos := 14
	len := 17
	v1 := 10000
	v2 , err := B(pos , len , v1)
	log(SEP)
	log("v2 : %v , err : %v  , pos:%v , len:%v , v1:%v\n" , v2 , err , pos , len , v1)
	if v2 != 10000 {
		t.Errorf("Failed.")
	}
}


func Test_5(t *testing.T){
	
}