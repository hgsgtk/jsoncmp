# JSON value comparison

[![Build Status](https://travis-ci.com/hgsgtk/jsoncmp.svg?branch=master)](https://travis-ci.com/hgsgtk/jsoncmp)
[![Go Report Card](https://goreportcard.com/badge/github.com/hgsgtk/jsoncmp)](https://goreportcard.com/report/github.com/hgsgtk/jsoncmp)
[![Coverage Status](https://coveralls.io/repos/github/hgsgtk/jsoncmp/badge.svg?branch=master)](https://coveralls.io/github/hgsgtk/jsoncmp?branch=master)
[![GoDoc](https://godoc.org/github.com/hgsgtk/jsoncmp?status.svg)](https://godoc.org/github.com/hgsgtk/jsoncmp)
[![Latest version](https://img.shields.io/github/tag/hgsgtk/jsoncmp.svg?label=latest%20version&style=flat)](https://github.com/hgsgtk/jsoncmp/releases)
[![license](https://img.shields.io/badge/license-MIT-4183c4.svg)](https://github.com/hgsgtk/jsoncmp/blob/master/LICENSE)


A Go JSON value comparison library for testing. It wraps [github.com/google/go-cmp][gocmp] function.
See the [GoDoc documentation][godoc] for more information.

[gocmp]: https://github.com/google/go-cmp
[godoc]: https://godoc.org/github.com/hgsgtk/jsoncmp

## Install

```
go get -u github.com/hgsgtk/jsoncmp
```

## Usage

```go
func TestDiff() {
	x := `
{
	"name": "Tom Bake",
	"age": 41
}`
	y := `
{
	"name": "Tom Bake",
	"age": 42
}
`
	var t stubT
	if diff := jsoncmp.Diff(x, y); diff != "" {
		t.Errorf("jsoncmp.Diff() got differs: (-got +want)\n%s", diff)
	}
}
```

## License

MIT - See [LICENSE][license] file

[license]: https://github.com/hgsgtk/jsoncmp/blob/master/LICENSE