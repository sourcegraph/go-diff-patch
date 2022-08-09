package godiffpatch

import (
	"flag"
	"os"
	"path"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var updateRegex *regexp.Regexp

func TestMain(m *testing.M) {
	// This allows us to provide an -update flag when running these tests.

	// Parse the flags to see if we need to update the golden files.
	update := flag.String("update", "", "test names matching this regex will have their golden files updated")
	flag.Parse()

	if *update != "" {
		updateRegex = regexp.MustCompile(*update)
	}

	// Actually run the tests.
	os.Exit(m.Run())
}

func TestGeneratePatch(t *testing.T) {
	for _, name := range []string{
		"PatchWithAddedLines",
		"PatchWithDeletedAddedUnchangedLines",
		"PatchWithDeletedLines",
		"PatchWithoutTrailingNewLine",
	} {
		t.Run(name, func(t *testing.T) {
			f := fixture{path.Join("fixtures", name)}

			patch := GeneratePatch("original", f.Original(t), f.Updated(t))
			f.AssertGolden(t, patch)
		})
	}
}

// fixture represents a single test directory with files named:
//
// - original
// - updated
// - golden
//
// golden will be updated if updateRegex is non-nil and matches the name of the
// test that invokes assertGolden.
type fixture struct {
	path string
}

func (f fixture) Original(t *testing.T) string {
	t.Helper()
	return string(f.readFile(t, "original"))
}

func (f fixture) Updated(t *testing.T) string {
	t.Helper()
	return string(f.readFile(t, "updated"))
}

func (f fixture) AssertGolden(t *testing.T, have string) {
	t.Helper()
	if updateRegex != nil && updateRegex.MatchString(t.Name()) {
		require.NoError(t, os.WriteFile(path.Join(f.path, "golden"), []byte(have), 0666))
		t.Logf("updated %s/golden", f.path)
	} else {
		want := f.readFile(t, "golden")
		assert.Equal(t, string(want), have)
	}
}

func (f fixture) readFile(t *testing.T, file string) []byte {
	t.Helper()

	path := path.Join(f.path, file)
	content, err := os.ReadFile(path)
	require.NoError(t, err)

	return content
}
