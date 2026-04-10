# go-pckstr

Package **pckstr** provides **packed-strings**, for the Go programming-language (golang).
A packed-strings is an alternative to `[]string` that can be compared using `==` and `!=` operators.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-pckstr

[![GoDoc](https://godoc.org/github.com/reiver/go-pckstr?status.svg)](https://godoc.org/github.com/reiver/go-pckstr)

## Example

Here is an example of creating a packed-strings:

```golang
import "github.com/reiver/go-pckstr"

packed := pckstr.SomeStrings("Once", "Twice", "Thrice", "Fource")
```

Here is an example of unpacking a packed-strings:

```golang
import "github.com/reiver/go-pckstr"

strings := packed.Strings()
// strings == []string{"Once", "Twice", "Thrice", "Fource"}
```

## Import

To import package **pckstr** use `import` code like the follownig:

```
import "github.com/reiver/go-pckstr"
```

## Installation

To install package **pckstr** do the following:

```
GOPROXY=direct go get github.com/reiver/go-pckstr
```

## Author

Package **pckstr** was written by [Charles Iliya Krempeaux](http://reiver.link)
