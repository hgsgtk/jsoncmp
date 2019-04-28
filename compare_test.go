package jsoncmp_test

import (
	"github.com/hgsgtk/jsoncmp"
	"testing"
)

type test struct {
	name string
	x, y string
	want string
}

func TestDiff(t *testing.T) {
	tests := compareTests(t)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jsoncmp.Diff(tt.x, tt.y); got != tt.want {
				t.Errorf("jsoncmp.Diff() got '%s', want '%s'", got, tt.want)
			}
		})
	}
}

func compareTests(t *testing.T) []test {
	t.Helper()

	return []test{
		{
			name: "EquivalentAndEqual",
			x: `
{
	"hoge": "huga",
	"piyo": "bar"
}`,
			y:    `{"hoge": "huga", "piyo": "bar"}`,
			want: "",
		},
		{
			name: "EquivalentAndNotEqual",
			x:    `{"hoge": "huga", "piyo": "bar"}`,
			y:    `{"piyo": "bar", "hoge": "huga"}`,
			want: "",
		},
		{
			name: "NotEquivalent",
			x:    `{"hoge": "huga", "piyo": "bar1"}`,
			y:    `{"piyo": "bar", "hoge": "huga"}`,
			want: `  string(Inverse(JSONcmp, map[string]interface{}{
    "hoge": string("huga"),
- 	"piyo": string("bar1"),
+ 	"piyo": string("bar"),
}))
`,
		},
	}
}
