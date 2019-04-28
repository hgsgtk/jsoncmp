package jsoncmp_test

import (
	"fmt"
	"github.com/hgsgtk/jsoncmp"
)

// Use Diff for comparing two JSON values and printing out human-redable errors.
func ExampleDiff() {
	x := `
{
	"name": "Tom Bake",
	"age": 41
}`
	y := `
{
	"name": "Tom Bake",
	"age": 42
}
`
	var t stubT
	if diff := jsoncmp.Diff(x, y); diff != "" {
		t.Errorf("jsoncmp.Diff() got differs: (-got +want)\n%s", diff)
	}
}

type stubT struct{}

func (t *stubT) Errorf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
