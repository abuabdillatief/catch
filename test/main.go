package main

import "github.com/abuabdillatief/catch"

type X struct {
	Name string
	Age int	
}


func main() {
	a := X{"Rendra", 34}
	catch.PrintStruct(a)
}