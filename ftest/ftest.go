package ftest

import (
	"fmt"
	"reflect"
	"testing"
)

// Test is data for assertion tests
type Test struct {
	ID     int
	Actual interface{}
	Wanted interface{}
}

// Assert runs assertion tests verifying equality of actual and wanted
func Assert(t *testing.T, tests []Test, f func(e string)) {
	for _, test := range tests {
		if test.Actual != test.Wanted {
			s := fmt.Sprintf("Failed Test[%d] => %v != %v ; %v wanted", test.ID, test.Actual, test.Wanted, test.Wanted)
			f(s)
		}
	}
}

// AssertDeep runs assertion tests verifying equality of actual and wanted
// tests for deep equality
func AssertDeep(t *testing.T, tests []Test, f func(e string)) {
	for _, test := range tests {
		if !reflect.DeepEqual(test.Actual, test.Wanted) {
			s := fmt.Sprintf("Failed [%d] => %v != %v ; %v wanted", test.ID, test.Actual, test.Wanted, test.Wanted)
			f(s)
		}
	}
}
