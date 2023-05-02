package main

import (
	"myapp/packageone"
)

var myVar = "this is my var in the package level"

func main() {

	var blockVar = "this is blockvar"

	packageone.PrintMe(myVar, blockVar)
}
