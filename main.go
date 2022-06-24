package main

import (
	"fmt"

	"github.com/rhiadc/fuzz-tests/functions"
	"github.com/rhiadc/fuzz-tests/overwrite"
)

func main() {
	input := "The quick brown fox jumped over the lazy dog"
	rev := functions.Reverse(input)
	doubleRev := functions.Reverse(rev)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q\n", rev)
	fmt.Printf("reversed again: %q\n", doubleRev)

	string := "0000\u00a00cd"
	newString := overwrite.OverwriteString(string, '\x01', 6)
	fmt.Println(newString)
}
