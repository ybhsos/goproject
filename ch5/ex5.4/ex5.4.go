package main

import "fmt"

func main() {

	str := "Hello\tGo\t\tworld\n\"Go\"is Awesome!\n"

	fmt.Print(str)
	fmt.Printf("%s", str)
	fmt.Printf("%q", str)

}
