package godiffpatch

import (
	"fmt"

	"github.com/sourcegraph/go-diff-patch/internal/diff"
	"github.com/sourcegraph/go-diff-patch/internal/diff/myers"
	"github.com/sourcegraph/go-diff-patch/internal/span"
)

// GeneratePatch generates a unified diff that is git-compatible and highlights the difference
// between originalContent and updatedContent.
func GeneratePatch(filename, originalContent, updatedContent string) string {
	uri := span.URI(filename)
	diffs, err := myers.ComputeEdits(uri, originalContent, updatedContent)
	if err != nil {
		// An error can never be returned here. It is always nil. The returned error is to satisfy the interface defined in `internal/diff/diff.go#L24`.
		// the diffs calculated using Eugene Myers algorithm.
		// Error is returned to satisfy the interface defined in `internal/diff/diff.go#L24`
		panic(err)
	}

	unifiedDiff := diff.ToUnified(
		"a/"+filename,
		"b/"+filename,
		originalContent,
		diffs,
	)

	return fmt.Sprint(unifiedDiff)
}
