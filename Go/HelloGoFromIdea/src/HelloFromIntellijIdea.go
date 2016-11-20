package main

import ("fmt"
	"mypackage"
)

func main() {
	fmt.Println("Hello")
	var i int = 3
	i = i * i
	fmt.Println(i)
	for j := 0; j < i; j++ {
		fmt.Println(j+j)
	}
	fmt.Println(mypackage.ReturnHello());


	// how I compile package
	// 1) move MyPackage.go into $GOPATH/src/mypackage folder
	// 2) insdide this folder, perform go install
	// 3) can use this package in my app!!!

}
