package main

import "fmt"
import "./gonext"

var pnum = 45

var print = fmt.Println

func main() {

	// var inum int
	// var fnum float64
	// var str string
	// var cnum complex128
	// var blogic bool

	// var inum int = 5
	// var fnum float64 = 6.3
	// var str string = "hello"
	// var cnum complex128 = 4+5i
	// var blogic bool = true

	// var inum = 5
	// var fnum = 6.3
	// var str  = "hello"
	// var cnum  = 4+5i
	// var blogic  = true

	inum := 5
	fnum := 6.3
	str := "hello"
	cnum := 4 + 5i
	blogic := true

	const pi = 3.141

	fmt.Println(inum, fnum, str, cnum, blogic)
	fmt.Printf("%T,%T,%T,%T,%T \n", inum, fnum, str, cnum, blogic)

	

	pnum = 70
	fmt.Println(pnum)

	abc()

	fmt.Println(checkNum)

	fmt.Println("@ gonext :", gonext.Pck01Num)

	fmt.Println(pi)

	print("This is not fmt.Println")

}

func abc(){

	fmt.Println("@ abc: ",pnum)
}
