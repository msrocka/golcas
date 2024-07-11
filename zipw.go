package golcas

import (
	"archive/zip"
	"encoding/json"
	"os"
)

type ZipWriter struct {
	file   *os.File
	writer *zip.Writer
}

func NewPackWriter(filePath string) (*ZipWriter, error) {
	file, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	writer := &ZipWriter{
		file:   file,
		writer: zip.NewWriter(file)}
	return writer, nil
}

func (w *ZipWriter) Close() error {
	err := w.writer.Close()
	if err != nil {
		return err
	}
	return w.file.Close()
}

func (w *ZipWriter) WriteActor(a *Actor) error {
	return w.write("actors/"+a.ID+".json", a)
}

func (w *ZipWriter) WriteCurrency(c *Currency) error {
	return w.write("currencies/"+c.ID+".json", c)
}

func (w *ZipWriter) WriteDQSystem(d *DQSystem) error {
	return w.write("dq_systems/"+d.ID+".json", d)
}

func (w *ZipWriter) WriteEpd(e *Epd) error {
	return w.write("epds/"+e.ID+".json", e)
}

func (w *ZipWriter) WriteFlow(f *Flow) error {
	return w.write("flows/"+f.ID+".json", f)
}

func (w *ZipWriter) WriteFlowProperty(f *FlowProperty) error {
	return w.write("flow_properties/"+f.ID+".json", f)
}

func (w *ZipWriter) WriteImpactCategory(l *ImpactCategory) error {
	return w.write("lcia_categories/"+l.ID+".json", l)
}

func (w *ZipWriter) WriteImpactMethod(l *ImpactMethod) error {
	return w.write("lcia_methods/"+l.ID+".json", l)
}

func (w *ZipWriter) WriteLocation(l *Location) error {
	return w.write("locations/"+l.ID+".json", l)
}

func (w *ZipWriter) WriteParameter(p *Parameter) error {
	return w.write("parameters/"+p.ID+".json", p)
}

func (w *ZipWriter) WriteProcess(p *Process) error {
	return w.write("processes/"+p.ID+".json", p)
}

func (w *ZipWriter) WriteProductSystem(p *ProductSystem) error {
	return w.write("product_systems/"+p.ID+".json", p)
}

func (w *ZipWriter) WriteProject(p *Project) error {
	return w.write("projects/"+p.ID+".json", p)
}

func (w *ZipWriter) WriteResult(r *Result) error {
	return w.write("results/"+r.ID+".json", r)
}

func (w *ZipWriter) WriteSocialIndicator(s *SocialIndicator) error {
	return w.write("social_indicators/"+s.ID+".json", s)
}

func (w *ZipWriter) WriteSource(s *Source) error {
	return w.write("sources/"+s.ID+".json", s)
}

func (w *ZipWriter) WriteUnitGroup(u *UnitGroup) error {
	return w.write("unit_groups/"+u.ID+".json", u)
}

func (w *ZipWriter) write(path string, val any) error {
	bytes, err := json.Marshal(val)
	if err != nil {
		return err
	}
	return w.Write(path, bytes)
}

func (w *ZipWriter) Write(path string, bytes []byte) error {
	writer, err := w.writer.Create(path)
	if err != nil {
		return err
	}
	_, err = writer.Write(bytes)
	return err
}
