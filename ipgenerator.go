


// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strconv"
)

var binString1 int64
var binString2 int64
var binString3 int64
var binString4 int64

func main() {
	total := 1000000 //4294967295

	for i := 0; i < total; i++ {
		elem := fmt.Sprintf("%032b\n", i)

		binString1, _ = strconv.ParseInt(elem[0:8], 2, 64)
		binString2, _ = strconv.ParseInt(elem[8:16], 2, 64)
		binString3, _ = strconv.ParseInt(elem[16:24], 2, 64)
		binString4, _ = strconv.ParseInt(elem[24:32], 2, 64)
		fmt.Println(binString1, ".", binString2, ".", binString3, ".", binString4)
	}
}
