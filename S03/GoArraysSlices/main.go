package main

import "fmt"

func main(){

	//Array

	// var inumArray = [...]int{10,20,30,40,50}
	// inumArray := [...]int{10,20,30,40,50}

	// inumArray[2] = 500

	// // ifloatArray := [2][2]float64{{1.2,4.3},{4.6,9.2}}

	// inumArray =append(inumArray,100)

	// fmt.Println(inumArray)

	//slice

	var inumSlice = []int{10,20,30,40,50}

	inumSlice = append(inumSlice,1000)

	fmt.Println(inumSlice[1:3])


}