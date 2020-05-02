# jsonnetdoc

Documentation parser for [Jsdoc](https://jsdoc.app/) style comments in Jsonnet.

This is not even close to covering the complete spec. Only descriptions, @param,
and @return. The initial motivation for this project was
[Grafonnet](https://github.com/grafana/grafonnet-lib). This project has used
this documentation style for quite a while without a parser. Choices made for
this program were influenced by pre-existing comments in that code base.

See [testdata/](./testdata) for examples.
