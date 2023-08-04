package modA

import (
	"fmt"
	"time"
)

func A(){
	fmt.Printf("Hello world.\n")
	time.Sleep(2 * time.Second)

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