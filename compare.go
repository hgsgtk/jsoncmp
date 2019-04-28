// Package jsoncmp is a Go JSON value comparison library for testing.
package jsoncmp

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-cmp/cmp"
)

// Diff returns a compared result of differences between two values.
// It returns an empty string if there is no differences, but else
// returns a report as string.
func Diff(x, y string) string {
	xform := cmp.Transformer("JSONcmp", func(s string) (m map[string]interface{}) {
		if err := json.Unmarshal([]byte(s), &m); err != nil {
			panic(fmt.Sprintf("json.Unmarshal(%s) got unexpected error %#v", s, err))
		}
		return m
	})
	opt := cmp.FilterPath(func(p cmp.Path) bool {
		for _, ps := range p {
			if tr, ok := ps.(cmp.Transform); ok && tr.Option() == xform {
				return false
			}
		}
		return true
	}, xform)
	return cmp.Diff(x, y, opt)
}
