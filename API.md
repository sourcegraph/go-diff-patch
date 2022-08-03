# Intro

What should the API look like?

- Compute will not be able to change the mode of the file
- 

```
--- a/rollup.config.ts
+++ b/rollup.config.ts
@@ -1,40 +1,39 @@
```

The above is probably the minimal headers needed for a patch to be `git-compatible` (can be applied with `git apply <path/to/patch>`). 
The `a` and `b` are important.

## Inputs

* filename
* original content
* updated content

## Output

* Unified diff (string)

## Unsupported

* No support for binary
