package jsoncmp_test

import (
	"github.com/hgsgtk/jsoncmp"
	"io/ioutil"
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
				t.Errorf("jsoncmp.Diff() got %s, want %s", got, tt.want)
			}
		})
	}
}

func compareTests(t *testing.T) []test {
	t.Helper()

	return []test{
		{
			name: "Compare",
			x:    getStringByFile(t, "testdata/A.json.golden"),
			y:    getStringByFile(t, "testdata/A.json.golden"),
			want: "",
		},
	}
}

func getStringByFile(t *testing.T, path string) string {
	t.Helper()

	bs, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatalf("ioutil.ReadAll(%s) got unexpected error %#v", path, err)
	}
	return string(bs)
}
