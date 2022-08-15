package godiffpatch

import (
	"fmt"

	"github.com/sourcegraph/go-diff-patch/internal/diff"
	"github.com/sourcegraph/go-diff-patch/internal/diff/myers"
	"github.com/sourcegraph/go-diff-patch/internal/span"
)

// TODO: add description
func GeneratePatch(filename, originalContent, updatedContent string) string {
	uri := span.URI(filename)
	diffs, err := myers.ComputeEdits(uri, originalContent, updatedContent)
	if err != nil {
		// Ideally, this would never happen because the error returned here is always nil.
		// Error is returned to satisfy the interface defined in `diff/diff.go#L24`
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
