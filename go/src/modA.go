package modA

import (
	"fmt"
	"time"
)

var timeSleep = time.Sleep

type myInterface interface{
	MyMethodA()string
	MyMethodB(int)float64
}
type sampleStruct struct{}

func A(){
	fmt.Printf("Hello world.\n")
	timeSleep(2 * time.Second)

}

func B () {
	go A () 
	fmt.Printf("Exiting B \n")
}

func C () {
	B()
	time.Sleep(5 * time.Second)
	fmt.Printf("Exiting C \n")
}

func (sampleStruct)myMethodA()(string){
	return "This is sampleStruct."
}

func (sampleStruct)myMethodB(a int)(ret float64){
	return float64(a*a)
}

func sampleFunction(m myInterface)(){
	m.MyMethodA()
	m.MyMethodB(5)
}