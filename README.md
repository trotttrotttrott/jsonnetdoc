# jsonnetdoc

Documentation parser for [JSDoc](https://jsdoc.app/) style comments in Jsonnet.

This is not even close to covering the complete spec. Only descriptions,
`@name`, `@param`, and `@return`. The initial motivation for this was
[Grafonnet](https://github.com/grafana/grafonnet-lib) which has used this
documentation style for quite a while without something to parse it.

See [testdata/](./testdata) for documentation examples.

## Usage

```
Usage:
  jsonnetdoc <input-file|dir> [flags]

Flags:
  -h, --help       help for jsonnetdoc
      --markdown   output markdown instead of JSON
```
