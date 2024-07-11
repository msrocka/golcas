# golcas

`golcas` is a Go package for writing reading and writing data in the 
[olca-schema format](https://github.com/GreenDelta/olca-schema)
(a JSON-[LD](https://json-ld.org/) based format for
[LCA](https://en.wikipedia.org/wiki/Life-cycle_assessment) data).

## Getting started
Install [Go](https://golang.org/) and run `go get` to retrieve this library:

```
$ go get github.com/msrocka/golcas...
```

### Reading and writing data
`golcas` has very fast methods for reading and writing data directly from
olca-schema packages, e.g.:

```go
import "github.com/msrocka/golcas"

r, err := golcas.NewPackReader("path/to/file.zip")
if err != nil {
    return err
}
err = r.EachFlow(func(f *golcas.Flow) bool {
    // do something with f
    return true  // return false, if you do not want to continue
})
```