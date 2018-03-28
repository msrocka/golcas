package ld

import (
	"archive/zip"
	"encoding/json"
	"io/ioutil"
)

// PackReader reads data sets from a zip file in the olca-schema
// package format.
type PackReader struct {
	reader *zip.ReadCloser
}

// NewPackReader creates a new PackReader
func NewPackReader(filePath string) (*PackReader, error) {
	reader, err := zip.OpenReader(filePath)
	return &PackReader{reader: reader}, err
}

// Close closes the PackReader
func (r *PackReader) Close() error {
	return r.reader.Close()
}

// GetActor reads the Actor with the given ID from the package
func (r *PackReader) GetActor(id string) (*Actor, error) {
	fname := "actors/" + id + ".json"
	for _, f := range r.reader.File {
		if f.Name != fname {
			continue
		}
		reader, err := f.Open()
		if err != nil {
			return nil, err
		}
		bytes, err := ioutil.ReadAll(reader)
		if err != nil {
			return nil, err
		}
		a := &Actor{}
		err = json.Unmarshal(bytes, a)
		return a, err
	}
	return nil, nil
}

// EachFile calls the given function for each file in the zip package. It stops
// when the function returns false or when there are no more files in the
// package.
func (r *PackReader) EachFile(fn func(f *ZipFile) bool) {
	files := r.reader.File
	for i := range files {
		zf := newZipFile(files[i])
		if !fn(zf) {
			break
		}
	}
}
