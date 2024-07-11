package golcas

import (
	"archive/zip"
	"io"
	"strings"
)

type ZipReader struct {
	reader *zip.ReadCloser
}

func NewZipReader(filePath string) (*ZipReader, error) {
	reader, err := zip.OpenReader(filePath)
	if err != nil {
		return nil, err
	}
	return &ZipReader{reader: reader}, nil
}

func (r *ZipReader) Close() error {
	return r.reader.Close()
}

func (r *ZipReader) ReadActor(id string) (*Actor, error) {
	return readPackModel(r, "actors/"+id+".json", ReadActor)
}

func (r *ZipReader) ReadCurrency(id string) (*Currency, error) {
	return readPackModel(r, "currencies/"+id+".json", ReadCurrency)
}

func (r *ZipReader) ReadDQSystem(id string) (*DQSystem, error) {
	return readPackModel(r, "dq_systems/"+id+".json", ReadDQSystem)
}

func (r *ZipReader) ReadEpd(id string) (*Epd, error) {
	return readPackModel(r, "epds/"+id+".json", ReadEpd)
}

func (r *ZipReader) ReadFlow(id string) (*Flow, error) {
	return readPackModel(r, "flows/"+id+".json", ReadFlow)
}

func (r *ZipReader) ReadFlowProperty(id string) (*FlowProperty, error) {
	return readPackModel(r, "flow_properties/"+id+".json", ReadFlowProperty)
}

func (r *ZipReader) ReadImpactCategory(id string) (*ImpactCategory, error) {
	return readPackModel(r, "lcia_categories/"+id+".json", ReadImpactCategory)
}

func (r *ZipReader) ReadImpactMethod(id string) (*ImpactMethod, error) {
	return readPackModel(r, "lcia_methods/"+id+".json", ReadImpactMethod)
}

func (r *ZipReader) ReadLocation(id string) (*Location, error) {
	return readPackModel(r, "locations/"+id+".json", ReadLocation)
}

func (r *ZipReader) ReadParameter(id string) (*Parameter, error) {
	return readPackModel(r, "parameters/"+id+".json", ReadParameter)
}

func (r *ZipReader) ReadProcess(id string) (*Process, error) {
	return readPackModel(r, "processes/"+id+".json", ReadProcess)
}

func (r *ZipReader) ReadProductSystem(id string) (*ProductSystem, error) {
	return readPackModel(r, "product_systems/"+id+".json", ReadProductSystem)
}

func (r *ZipReader) ReadProject(id string) (*Project, error) {
	return readPackModel(r, "projects/"+id+".json", ReadProject)
}

func (r *ZipReader) ReadResult(id string) (*Result, error) {
	return readPackModel(r, "results/"+id+".json", ReadResult)
}

func (r *ZipReader) ReadSocialIndicator(id string) (*SocialIndicator, error) {
	return readPackModel(r, "social_indicators/"+id+".json", ReadSocialIndicator)
}

func (r *ZipReader) ReadSource(id string) (*Source, error) {
	return readPackModel(r, "sources/"+id+".json", ReadSource)
}

func (r *ZipReader) ReadUnitGroup(id string) (*UnitGroup, error) {
	return readPackModel(r, "unit_groups/"+id+".json", ReadUnitGroup)
}

func (r *ZipReader) EachActor(fn func(*Actor) bool) error {
	return eachPackModel(r, "actors/", ReadActor, fn)
}

func (r *ZipReader) EachCurrency(fn func(*Currency) bool) error {
	return eachPackModel(r, "currencies/", ReadCurrency, fn)
}

func (r *ZipReader) EachDQSystem(fn func(*DQSystem) bool) error {
	return eachPackModel(r, "dq_systems/", ReadDQSystem, fn)
}

func (r *ZipReader) EachEpd(fn func(*Epd) bool) error {
	return eachPackModel(r, "epds/", ReadEpd, fn)
}

func (r *ZipReader) EachFlow(fn func(*Flow) bool) error {
	return eachPackModel(r, "flows/", ReadFlow, fn)
}

func (r *ZipReader) EachFlowProperty(fn func(*FlowProperty) bool) error {
	return eachPackModel(r, "flow_properties/", ReadFlowProperty, fn)
}

func (r *ZipReader) EachImpactCategory(fn func(*ImpactCategory) bool) error {
	return eachPackModel(r, "lcia_categories/", ReadImpactCategory, fn)
}

func (r *ZipReader) EachImpactMethod(fn func(*ImpactMethod) bool) error {
	return eachPackModel(r, "lcia_methods/", ReadImpactMethod, fn)
}

func (r *ZipReader) EachLocation(fn func(*Location) bool) error {
	return eachPackModel(r, "locations/", ReadLocation, fn)
}

func (r *ZipReader) EachParameter(fn func(*Parameter) bool) error {
	return eachPackModel(r, "parameters/", ReadParameter, fn)
}

func (r *ZipReader) EachProcess(fn func(*Process) bool) error {
	return eachPackModel(r, "processes/", ReadProcess, fn)
}

func (r *ZipReader) EachProductSystem(fn func(*ProductSystem) bool) error {
	return eachPackModel(r, "product_systems/", ReadProductSystem, fn)
}

func (r *ZipReader) EachProject(fn func(*Project) bool) error {
	return eachPackModel(r, "projects/", ReadProject, fn)
}

func (r *ZipReader) EachResult(fn func(*Result) bool) error {
	return eachPackModel(r, "results/", ReadResult, fn)
}

func (r *ZipReader) EachSocialIndicator(fn func(*SocialIndicator) bool) error {
	return eachPackModel(r, "social_indicators/", ReadSocialIndicator, fn)
}

func (r *ZipReader) EachSource(fn func(*Source) bool) error {
	return eachPackModel(r, "sources/", ReadSource, fn)
}

func (r *ZipReader) EachUnitGroup(fn func(*UnitGroup) bool) error {
	return eachPackModel(r, "unit_groups/", ReadUnitGroup, fn)
}

func readPackModel[T any](
	r *ZipReader, entry string, f func([]byte) (*T, error)) (*T, error) {
	e, err := r.reader.Open(entry)
	if err != nil {
		return nil, err
	}
	defer e.Close()
	bytes, err := io.ReadAll(e)
	if err != nil {
		return nil, err
	}
	return f(bytes)
}

func eachPackModel[T any](
	r *ZipReader,
	folder string,
	reader func([]byte) (*T, error),
	handler func(*T) bool) error {

	files := r.reader.File
	for i := range files {
		file := files[i]
		if file.FileInfo().IsDir() || !strings.HasPrefix(file.Name, folder) {
			continue
		}

		data, err := readZipEntry(file)
		if err != nil {
			return err
		}

		model, err := reader(data)
		if err != nil {
			return err
		}

		if !handler(model) {
			break
		}
	}
	return nil
}

func readZipEntry(f *zip.File) ([]byte, error) {
	reader, err := f.Open()
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	return io.ReadAll(reader)
}
