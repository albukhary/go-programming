/*
package main

import "fmt"

func main() {
	fmt.Println("Hello everyone, this is the most awesome class in the entire world for having fun and learning the GO programming language.")

	foo()

	fmt.Println("something more")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()
}

func foo() {
	fmt.Println("I'm in foo.")
}

func bar() {
	fmt.Println("and then we exited")
}
*/
// control flow:
// (1) sequence
// (2) loop; iterative
// (3) conditional

package main

import (
	"fmt"
)

func main(){
	fmt.Println("Assalamu alaykum guys! This a totally awesome course to follow. May Allah make our journey easy! Ameen")
	foo()
	fmt.Println("Something more...")
	for i:=0;i<10;i++{
		if i%2==0{
			fmt.Printf("%v is an even number bro., learn math\n",i)
		}
	}
	bar()
}
func foo(){
	fmt.Println("I am in foo() baby... ")

}
func bar(){
	fmt.Println("Exiting main()...")
}