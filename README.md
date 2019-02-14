# ld

[![Build Status](https://travis-ci.org/msrocka/ld.svg?branch=master)](https://travis-ci.org/msrocka/ld)

`ld` is a small Go package for writing reading and writing data in the 
[olca-schema format](https://github.com/GreenDelta/olca-schema) (a JSON-**LD**
based format for [LCA](https://en.wikipedia.org/wiki/Life-cycle_assessment) data).

## Getting started
Install [Go](https://golang.org/) and run `go get` to retrieve this library:

```
$ go get github.com/msrocka/ld...
```

### Reading and writing data
`ld` has very fast methods for reading and writing data directly from
olca-schema packages, e.g.:

```go
import "github.com/msrocka/ld"

r, err := ld.NewPackReader("path/to/file.zip")
if err != nil {
    return err
}
err = r.EachFlow(func(f *ld.Flow) bool {
    // do something with f
    return true  // return false, if you do not want to continue
})
```