package main

import (
	"myapp/packageone"
)

var myVar = "this is my var"

func main() {

	var blockVar = "this is blockvar"

	packageone.PrintMe(myVar, blockVar)
}
