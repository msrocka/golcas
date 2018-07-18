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
func (f *ZipFile) ReadCategory() (*Category, error) {
	data, err := f.Read()
	if err != nil {
		return nil, err
	}
	return ReadCategory(data)
}

// ReadActor reads a Actor from the zip file
func (f *ZipFile) ReadActor() (*Actor, error) {
	data, err := f.Read()
	if err != nil {
		return nil, err
	}
	return ReadActor(data)
}

// ReadSource reads a Source from the zip file
func (f *ZipFile) ReadSource() (*Source, error) {
	data, err := f.Read()
	if err != nil {
		return nil, err
	}
	return ReadSource(data)
}

// ReadParameter reads a Parameter from the zip file
func (f *ZipFile) ReadParameter() (*Parameter, error) {
	data, err := f.Read()
	if err != nil {
		return nil, err
	}
	return ReadParameter(data)
}

// ReadSocialIndicator reads a SocialIndicator from the zip file
func (f *ZipFile) ReadSocialIndicator() (*SocialIndicator, error) {
	data, err := f.Read()
	if err != nil {
		return nil, err
	}
	return ReadSocialIndicator(data)
}

// ReadLocation reads a Location from the zip file
func (f *ZipFile) ReadLocation() (*Location, error) {
	data, err := f.Read()
	if err != nil {
		return nil, err
	}
	return ReadLocation(data)
}

// ReadUnitGroup reads a UnitGroup from the zip file
func (f *ZipFile) ReadUnitGroup() (*UnitGroup, error) {
	data, err := f.Read()
	if err != nil {
		return nil, err
	}
	return ReadUnitGroup(data)
}

// ReadFlowProperty reads a FlowProperty from the zip file
func (f *ZipFile) ReadFlowProperty() (*FlowProperty, error) {
	data, err := f.Read()
	if err != nil {
		return nil, err
	}
	return ReadFlowProperty(data)
}

// ReadFlow reads a Flow from the zip file
func (f *ZipFile) ReadFlow() (*Flow, error) {
	data, err := f.Read()
	if err != nil {
		return nil, err
	}
	return ReadFlow(data)
}

// ReadProcess reads a Process from the zip file
func (f *ZipFile) ReadProcess() (*Process, error) {
	data, err := f.Read()
	if err != nil {
		return nil, err
	}
	return ReadProcess(data)
}

// ReadImpactCategory reads a ImpactCategory from the zip file
func (f *ZipFile) ReadImpactCategory() (*ImpactCategory, error) {
	data, err := f.Read()
	if err != nil {
		return nil, err
	}
	return ReadImpactCategory(data)
}

// ReadImpactMethod reads a ImpactMethod from the zip file
func (f *ZipFile) ReadImpactMethod() (*ImpactMethod, error) {
	data, err := f.Read()
	if err != nil {
		return nil, err
	}
	return ReadImpactMethod(data)
}
