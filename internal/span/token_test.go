// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package span

import (
	"fmt"
	"go/token"
	"path"
	"testing"
)

var testdata = []struct {
	uri     string
	content []byte
}{
	{"/a.go", []byte(`
// file a.go
package test
`)},
	{"/b.go", []byte(`
//
//
// file b.go
package test`)},
	{"/c.go", []byte(`
// file c.go
package test`)},
}

var tokenTests = []Span{
	New(URIFromPath("/a.go"), NewPoint(1, 1, 0), Point{}),
	New(URIFromPath("/a.go"), NewPoint(3, 7, 20), NewPoint(3, 7, 20)),
	New(URIFromPath("/b.go"), NewPoint(4, 9, 15), NewPoint(4, 13, 19)),
	New(URIFromPath("/c.go"), NewPoint(4, 1, 26), Point{}),
}

func TestToken(t *testing.T) {
	fset := token.NewFileSet()
	files := map[URI]*token.File{}
	for _, f := range testdata {
		file := fset.AddFile(f.uri, -1, len(f.content))
		file.SetLinesForContent(f.content)
		files[URIFromPath(f.uri)] = file
	}
	for _, test := range tokenTests {
		f := files[test.URI()]
		t.Run(path.Base(f.Name()), func(t *testing.T) {
			checkToken(t, f, New(
				test.URI(),
				NewPoint(test.Start().Line(), test.Start().Column(), 0),
				NewPoint(test.End().Line(), test.End().Column(), 0),
			), test)
			checkToken(t, f, New(
				test.URI(),
				NewPoint(0, 0, test.Start().Offset()),
				NewPoint(0, 0, test.End().Offset()),
			), test)
		})
	}
}

func checkToken(t *testing.T, f *token.File, in, expect Span) {
	rng, err := in.Range(f)
	if err != nil {
		t.Error(err)
	}
	gotLoc, err := rng.Span()
	if err != nil {
		t.Error(err)
	}
	expected := fmt.Sprintf("%+v", expect)
	got := fmt.Sprintf("%+v", gotLoc)
	if expected != got {
		t.Errorf("For %v expected %q got %q", in, expected, got)
	}
}
