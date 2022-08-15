// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !windows
// +build !windows

package span

import "testing"

// TestURI tests the conversion between URIs and filenames. The test cases
// include Windows-style URIs and filepaths, but we avoid having OS-specific
// tests by using only forward slashes, assuming that the standard library
// functions filepath.ToSlash and filepath.FromSlash do not need testing.
func TestURIFromPath(t *testing.T) {
	for _, test := range []struct {
		path, wantFile string
		wantURI        URI
	}{
		{
			path:     ``,
			wantFile: ``,
			wantURI:  URI(""),
		},
		{
			path:     `C:/Windows/System32`,
			wantFile: `C:/Windows/System32`,
			wantURI:  URI("file:///C:/Windows/System32"),
		},
		{
			path:     `C:/Go/src/bob.go`,
			wantFile: `C:/Go/src/bob.go`,
			wantURI:  URI("file:///C:/Go/src/bob.go"),
		},
		{
			path:     `c:/Go/src/bob.go`,
			wantFile: `C:/Go/src/bob.go`,
			wantURI:  URI("file:///C:/Go/src/bob.go"),
		},
		{
			path:     `/path/to/dir`,
			wantFile: `/path/to/dir`,
			wantURI:  URI("file:///path/to/dir"),
		},
		{
			path:     `/a/b/c/src/bob.go`,
			wantFile: `/a/b/c/src/bob.go`,
			wantURI:  URI("file:///a/b/c/src/bob.go"),
		},
		{
			path:     `c:/Go/src/bob george/george/george.go`,
			wantFile: `C:/Go/src/bob george/george/george.go`,
			wantURI:  URI("file:///C:/Go/src/bob%20george/george/george.go"),
		},
	} {
		got := URIFromPath(test.path)
		if got != test.wantURI {
			t.Errorf("URIFromPath(%q): got %q, expected %q", test.path, got, test.wantURI)
		}
		gotFilename := got.Filename()
		if gotFilename != test.wantFile {
			t.Errorf("Filename(%q): got %q, expected %q", got, gotFilename, test.wantFile)
		}
	}
}

func TestURIFromURI(t *testing.T) {
	for _, test := range []struct {
		inputURI, wantFile string
		wantURI            URI
	}{
		{
			inputURI: `file:///c:/Go/src/bob%20george/george/george.go`,
			wantFile: `C:/Go/src/bob george/george/george.go`,
			wantURI:  URI("file:///C:/Go/src/bob%20george/george/george.go"),
		},
		{
			inputURI: `file:///C%3A/Go/src/bob%20george/george/george.go`,
			wantFile: `C:/Go/src/bob george/george/george.go`,
			wantURI:  URI("file:///C:/Go/src/bob%20george/george/george.go"),
		},
		{
			inputURI: `file:///path/to/%25p%25ercent%25/per%25cent.go`,
			wantFile: `/path/to/%p%ercent%/per%cent.go`,
			wantURI:  URI(`file:///path/to/%25p%25ercent%25/per%25cent.go`),
		},
		{
			inputURI: `file:///C%3A/`,
			wantFile: `C:/`,
			wantURI:  URI(`file:///C:/`),
		},
		{
			inputURI: `file:///`,
			wantFile: `/`,
			wantURI:  URI(`file:///`),
		},
		{
			inputURI: `file://wsl%24/Ubuntu/home/wdcui/repo/VMEnclaves/cvm-runtime`,
			wantFile: `/wsl$/Ubuntu/home/wdcui/repo/VMEnclaves/cvm-runtime`,
			wantURI:  URI(`file:///wsl$/Ubuntu/home/wdcui/repo/VMEnclaves/cvm-runtime`),
		},
	} {
		got := URIFromURI(test.inputURI)
		if got != test.wantURI {
			t.Errorf("NewURI(%q): got %q, expected %q", test.inputURI, got, test.wantURI)
		}
		gotFilename := got.Filename()
		if gotFilename != test.wantFile {
			t.Errorf("Filename(%q): got %q, expected %q", got, gotFilename, test.wantFile)
		}
	}
}
