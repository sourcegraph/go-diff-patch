package example

import (
	"os"

	godiffpatch "github.com/sourcegraph/go-diff-patch"
)

func main() {
	patch := godiffpatch.GeneratePatch("test.txt", original, updated)
	err := os.WriteFile("test.txt", []byte(patch), 0644)
	if err != nil {
		panic(err)
	}
}

const original = ``

const updated = `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.`
