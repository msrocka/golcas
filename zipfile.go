package ld

import (
	"archive/zip"
	"io/ioutil"
)

// ZipFile embedds the type `File` from the `archive/zip` package and provides
// additional ILCD specific methods.
type ZipFile struct {
	f *zip.File
}

// newZipFile initializes a new ZipFile from the given archive file.
func newZipFile(f *zip.File) *ZipFile {
	return &ZipFile{f}
}

// Path returns the path of the zip file within the zip package.
func (f *ZipFile) Path() string {
	return f.f.Name
}

// Reads the decompressed data from the zip file.
func (f *ZipFile) Read() ([]byte, error) {
	reader, err := f.f.Open()
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	if err := reader.Close(); err != nil {
		return nil, err
	}
	return data, nil
}

// ReadCategory reads a Category from the zip file
func (f *ZipFile) ReadCategory(data []byte) (*Category, error) {
	data, err := f.Read()
	if err != nil {
		return nil, err
	}
	return ReadCategory(data)
}

// ReadActor reads a Actor from the zip file
func (f *ZipFile) ReadActor(data []byte) (*Actor, error) {
	data, err := f.Read()
	if err != nil {
		return nil, err
	}
	return ReadActor(data)
}

// ReadSource reads a Source from the zip file
func (f *ZipFile) ReadSource(data []byte) (*Source, error) {
	data, err := f.Read()
	if err != nil {
		return nil, err
	}
	return ReadSource(data)
}

// ReadUnitGroup reads a UnitGroup from the zip file
func (f *ZipFile) ReadUnitGroup(data []byte) (*UnitGroup, error) {
	data, err := f.Read()
	if err != nil {
		return nil, err
	}
	return ReadUnitGroup(data)
}

// ReadFlowProperty reads a FlowProperty from the zip file
func (f *ZipFile) ReadFlowProperty(data []byte) (*FlowProperty, error) {
	data, err := f.Read()
	if err != nil {
		return nil, err
	}
	return ReadFlowProperty(data)
}

// ReadFlow reads a Flow from the zip file
func (f *ZipFile) ReadFlow(data []byte) (*Flow, error) {
	data, err := f.Read()
	if err != nil {
		return nil, err
	}
	return ReadFlow(data)
}

// ReadProcess reads a Process from the zip file
func (f *ZipFile) ReadProcess(data []byte) (*Process, error) {
	data, err := f.Read()
	if err != nil {
		return nil, err
	}
	return ReadProcess(data)
}

// ReadImpactCategory reads a ImpactCategory from the zip file
func (f *ZipFile) ReadImpactCategory(data []byte) (*ImpactCategory, error) {
	data, err := f.Read()
	if err != nil {
		return nil, err
	}
	return ReadImpactCategory(data)
}

// ReadImpactMethod reads a ImpactMethod from the zip file
func (f *ZipFile) ReadImpactMethod(data []byte) (*ImpactMethod, error) {
	data, err := f.Read()
	if err != nil {
		return nil, err
	}
	return ReadImpactMethod(data)
}
