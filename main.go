package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/BolajiOlajide/go-tools/diff/myers"
	"github.com/BolajiOlajide/go-tools/span"
)

func main() {
	uri := span.URI("rollup.config.ts")
	f, err := os.Create("generated.patch")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	diffs, err := myers.ComputeEdits(uri, originalContent, updatedContent)
	// diffs, err := gdiff.NComputeEdits(uri, originalContent, updatedContent)
	if err != nil {
		panic(err)
	}

	splittedString := strings.Split(originalContent, "\n")

	writeHeaders(f, Meta{
		FileName:  uri,
		FileMode:  100644,
		SourceSHA: "e247121",
		TargetSHA: "1df8fb8",
	})

	for i := 0; i < len(diffs); i++ {
		diff := diffs[i]
		startPoint := diff.Span.Start()
		endPoint := diff.Span.End()

		sl := startPoint.Line()
		el := endPoint.Line()

		fmt.Println(sl, el)

		if sl == 1 {
			fmt.Println(splittedString[sl-1 : el])
		}

		if i < len(splittedString) && sl != el {
			for _, oldLine := range splittedString[sl-1 : el] {
				// fmt.Pri
				f.WriteString(fmt.Sprintf("-%s\n", oldLine))
			}
		}

		// if i < len(splittedString) {
		// 	oldLine := splittedString[i]
		// 	f.WriteString(fmt.Sprintf("-%s\n", oldLine))
		// }
		if i == 0 {
			continue
		}
		f.WriteString(fmt.Sprintf("+%s", diff.NewText))
		if diff.NewText == "" {
			f.WriteString("\n")
		}
	}
}

type Meta struct {
	FileName  span.URI
	FileMode  os.FileMode
	SourceSHA string
	TargetSHA string
	// How do we calculate the hunks :(
}

func writeHeaders(f *os.File, m Meta) {
	//   diff --git a/rollup.config.ts b/rollup.config.ts
	// index e247121..1df8fb8 100644
	// --- a/rollup.config.ts
	// +++ b/rollup.config.ts
	// @@ -1,40 +1,39 @@
	f.WriteString(fmt.Sprintf("diff --git a/%s b/%s\n", m.FileName, m.FileName))
	f.WriteString(fmt.Sprintf("index %s..%s %d\n", m.SourceSHA, m.TargetSHA, m.FileMode))
	f.WriteString(fmt.Sprintf("--- a/%s\n", m.FileName))
	f.WriteString(fmt.Sprintf("+++ b/%s\n", m.FileName))
}

const originalContent string = `import babel from '@rollup/plugin-babel';
import { nodeResolve } from '@rollup/plugin-node-resolve';
import { terser } from 'rollup-plugin-terser';

import pkg from './package.json';


const extensions = ['.ts'];

export default {
  input: 'src/index.ts',
  plugins: [
    nodeResolve({ extensions }),
    babel({
      extensions,
      babelHelpers: 'bundled',
      exclude: 'node_modules/**'
    }),
    terser()
  ],
  output: [
    {
      file: pkg.main,
      format: 'cjs',
      exports: 'auto',
      sourcemap: true
    },
    {
      file: pkg.module,
      format: 'es',
      sourcemap: true
    },
    {
      name: pkg.name,
      file: pkg.umd,
      format: 'umd',
      sourcemap: true
    }
  ]
};
`

const updatedContent string = `import babel from '@rollup/plugin-babel'
import { nodeResolve } from '@rollup/plugin-node-resolve'
import { terser } from 'rollup-plugin-terser'

import pkg from './package.json'

const extensions = ['.ts']

export default {
    input: 'src/index.ts',
    plugins: [
        nodeResolve({ extensions }),
        babel({
            extensions,
            babelHelpers: 'bundled',
            exclude: 'node_modules/**',
        }),
        terser(),
    ],
    output: [
        {
            file: pkg.main,
            format: 'cjs',
            exports: 'auto',
            sourcemap: true,
        },
        {
            file: pkg.module,
            format: 'es',
            sourcemap: true,
        },
        {
            name: pkg.name,
            file: pkg.umd,
            format: 'umd',
            sourcemap: true,
        },
    ],
}
`
