package main

import "fmt"

/*
func main() {
	n, _ := fmt.Println("Hello, playground", 42, true)
	fmt.Println(n)
}
*/
func main(){
	n, _ := myPrint("Yeah bro I do not need fmt anymore...")
	myPrint(n," bytes printed")
}
func myPrint(a ... interface{})( int, error){
	return fmt.Println(a)
}
