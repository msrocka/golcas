# golcas

`golcas` is a Go package for writing reading and writing data in the [openLCA schema](http://greendelta.github.io/olca-schema/). Here is a small example:

```go
w, err := golcas.NewZipWriter("path/to/file.zip")

// write a flow
err := w.WriteFlow(&golcas.Flow{
  ID: "123abc",
  Name: "Carbon dioxide",
  // ...
})

r, err := golcas.NewZipReader("path/to/file.zip")

// read a single flow
flow, err := r.ReadFlow("123abc")

// read all flows
err := r.EachFlow(func(flow *golcas.Flow) bool {
  // do something with the flwo
  return true
})
```
