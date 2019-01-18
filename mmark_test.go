package markdown

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/ToasterKTN/markdown/parser"
)

func TestMmark(t *testing.T) {
	testfile := filepath.Join("testdata", "mmark.test")

	data, err := ioutil.ReadFile(testfile)
	if err != nil {
		t.Fatalf("failed to open file %q: %s", testfile, err)
	}

	testdata := bytes.Split(data, []byte("+++\n"))
	if len(testdata)%2 != 0 {
		t.Fatalf("odd test tuples: %d", len(testdata))
	}

	ext := parser.CommonExtensions | parser.Attributes | parser.OrderedListStart | parser.SuperSubscript | parser.Mmark
	for i := 0; i < len(testdata); i += 2 {
		p := parser.NewWithExtensions(ext)

		input := testdata[i]
		want := testdata[i+1]

		got := ToHTML([]byte(input), p, nil)

		if bytes.Compare(got, want) != 0 {
			t.Errorf("want (%d bytes) %s, got (%d bytes) %s, for input %q", len(want), want, len(got), got, input)
		}
	}
}
