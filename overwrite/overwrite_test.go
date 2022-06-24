package overwrite_test

import (
	"testing"

	"github.com/rhiadc/fuzz-tests/overwrite"
)

//go test -fuzz FuzzBasicOverwriteString
//go test -fuzz FuzzBasicOverwriteString/83d4198b9f662bca5d4b0090ab9ff0e91ced8716a49a76198baa0c8c18826ad2
func FuzzBasicOverwriteString(f *testing.F) {
	f.Fuzz(func(t *testing.T, str string, value rune, n int) {
		overwrite.OverwriteString(str, value, n)
	})
}
