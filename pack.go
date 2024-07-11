package golcas

import (
	"archive/zip"
	"encoding/json"
	"os"
)

// PackWriter writes data sets to a zip file in the olca-schema
// package format.
type PackWriter struct {
	file   *os.File
	writer *zip.Writer
}

// NewPackWriter creates a new PackWriter
func NewPackWriter(filePath string) (*PackWriter, error) {
	file, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	writer := &PackWriter{
		file:   file,
		writer: zip.NewWriter(file)}
	return writer, nil
}

// Close closes the PackWriter
func (w *PackWriter) Close() error {
	err := w.writer.Close()
	if err != nil {
		return err
	}
	return w.file.Close()
}

// PutActor writes the given actor to the package
func (w *PackWriter) PutActor(a *Actor) error {
	return w.put("actors/"+a.ID+".json", a)
}

// PutSource writes the given source to the package
func (w *PackWriter) PutSource(s *Source) error {
	return w.put("sources/"+s.ID+".json", s)
}

// PutUnitGroup writes the given unit group to the package
func (w *PackWriter) PutUnitGroup(ug *UnitGroup) error {
	return w.put("unit_groups/"+ug.ID+".json", ug)
}

// PutFlowProperty writes the given flow property to the package
func (w *PackWriter) PutFlowProperty(fp *FlowProperty) error {
	return w.put("flow_properties/"+fp.ID+".json", fp)
}

// PutFlow writes the given flow to the package
func (w *PackWriter) PutFlow(f *Flow) error {
	return w.put("flows/"+f.ID+".json", f)
}

// PutProcess writes the given process to the package
func (w *PackWriter) PutProcess(p *Process) error {
	return w.put("processes/"+p.ID+".json", p)
}

// PutSocialIndicator writes the given social indicator to the package
func (w *PackWriter) PutSocialIndicator(si *SocialIndicator) error {
	return w.put("social_indicators/"+si.ID+".json", si)
}

// PutImpactMethod writes the given LCIA method to the package
func (w *PackWriter) PutImpactMethod(im *ImpactMethod) error {
	return w.put("lcia_methods/"+im.ID+".json", im)
}

// PutImpactCategory writes the given LCIA category to the package
func (w *PackWriter) PutImpactCategory(ic *ImpactCategory) error {
	return w.put("lcia_categories/"+ic.ID+".json", ic)
}

// PutProductSystem writes the given product system to the package
func (w *PackWriter) PutProductSystem(ps *ProductSystem) error {
	return w.put("product_systems/"+ps.ID+".json", ps)
}

func (w *PackWriter) put(path string, val interface{}) error {
	bytes, err := json.Marshal(val)
	if err != nil {
		return err
	}
	return w.Put(path, bytes)
}

// Put writes the given bytes to the given path
func (w *PackWriter) Put(path string, bytes []byte) error {
	writer, err := w.writer.Create(path)
	if err != nil {
		return err
	}
	_, err = writer.Write(bytes)
	return err
}
