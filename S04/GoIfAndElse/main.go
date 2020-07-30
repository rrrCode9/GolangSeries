package main

import (

	"fmt"
	// "strings"
	"time"
)

func main(){

	
	// Ex-01: a simple if statement---------------------------------
	
	

	// if auth := true;auth{

	// 	fmt.Println("Hello User")
		 
	// }


	











	// EX-02: if else if-----------------------------------------


	// hash := "gshsjiksmHell7ondhyuyG8ohst8787s"

	// isHello := strings.Contains(hash,"Hello")
	
	// isGo := strings.Contains(hash,"Go")

	// if isHello && isGo{

	// 	fmt.Println("contains: Hello and Go")

	// }else if isHello || isGo{

	// 	fmt.Println("contains: either Hello or Go")

	// }else{							// not in next line

	// 	fmt.Println("Contains: Neither Hello nor Go")
	// }


























	////Challenge-----------------------------------------------------

	t1 := time.Now()

	t2 := time.Now()

	if t2.Sub(t1).Nanoseconds()==0{

		fmt.Println("t1 is equal to t2")
	}else{

		fmt.Println("t1 is not equal to t2")
	}

	// //what is the answer? write in comment section below.











}