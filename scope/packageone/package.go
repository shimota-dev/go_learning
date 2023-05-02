package packageone

import (
	"fmt"
)

var PackageVar = "This is packagevar in the packageone"

func PrintMe(s1, s2 string) {
	fmt.Println(s1, s2, PackageVar)

}
