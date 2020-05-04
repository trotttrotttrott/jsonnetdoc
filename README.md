# jsonnetdoc

Documentation parser for [JSDoc](https://jsdoc.app/) style comments in Jsonnet.

This is not even close to covering the complete spec. Only descriptions,
`@name`, `@param`, and `@return`. The initial motivation for this was
[Grafonnet](https://github.com/grafana/grafonnet-lib) which has used this
documentation style for quite a while without something to parse it.

See [testdata/](./testdata) for documentation examples.

## Installation

No release binaries are maintained at this time. but Docker images are:
https://hub.docker.com/repository/docker/trotttrotttrott/jsonnetdoc. If you
don't want to use Docker, clone the repo and use one of the usual Go
installation methods like `go install`.

## Usage

Expects a single argument which should be a path to your Jsonnet files. By
default, it will output a JSON representation of your documentation. If you
pass, the `--markdown` flag, it will instead output Markdown.

```
Usage:
  jsonnetdoc <input-file|dir> [flags]

Flags:
  -h, --help       help for jsonnetdoc
      --markdown   output markdown instead of JSON
```

Example:

```
jsonnetdoc testdata --markdown
```

Same command using Docker:

```
docker run --rm \
  -v $PWD:$PWD \
  -w $PWD \
  trotttrotttrott/jsonnetdoc \
  testdata --markdown
```
