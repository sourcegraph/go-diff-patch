# go-diff-patch 

[![Go CI](https://github.com/sourcegraph/go-diff-patch/actions/workflows/go-ci.yml/badge.svg)](https://github.com/sourcegraph/go-diff-patch/actions/workflows/go-ci.yml)
[![lsif-go](https://github.com/sourcegraph/go-diff-patch/actions/workflows/lsif-go.yml/badge.svg)](https://github.com/sourcegraph/go-diff-patch/actions/workflows/lsif-go.yml)

A fork of [Go tools](golang.org/x/tools). Go-diff-patch is an utility library that is used to generate git-compatible patches that are appliable to any repository.

## Installation

`go-diff-patch` can be installed using the `go get` command as shown below:

```sh
go get github.com/sourcegraph/go-diff-patch
```

### Usage

The library exports a function `GeneratePatch` which computes edits between the original and updated file contents and produces a unified diff for them as a string. This patch can then be applied to a valid git repository containing the specified file(s).

An example can be found in the [`examples/` directory](https://sourcegraph.com/github.com/sourcegraph/go-diff-patch/-/blob/examples/example.go?L1%3A1-20%3A1=).
