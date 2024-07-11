package golcas

import (
	"archive/zip"
	"io"
)

type PackReader struct {
	reader *zip.ReadCloser
}

func NewPackReader(filePath string) (*PackReader, error) {
	reader, err := zip.OpenReader(filePath)
	if err != nil {
		return nil, err
	}
	return &PackReader{reader: reader}, nil
}

func (r *PackReader) Close() error {
	return r.reader.Close()
}

func (r *PackReader) ReadActor(id string) (*Actor, error) {
	f, err := r.reader.Open("actors/" + id + ".json")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	bytes, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return ReadActor(bytes)
}

// EachFile calls the given function for each file in the zip package. It stops
// when the function returns false or when there are no more files in the
// package.
func (r *PackReader) EachFile(fn func(f *ZipFile) bool) {
	files := r.reader.File
	for i := range files {
		file := files[i]
		if file.FileInfo().IsDir() {
			continue
		}
		zf := newZipFile(file)
		if !fn(zf) {
			break
		}
	}
}

// EachSource iterates over each `Source` data set in the package unless
// the given handler returns false.
func (r *PackReader) EachSource(fn func(*Source) bool) error {
	var gerr error
	r.EachFile(func(f *ZipFile) bool {
		if !IsSourcePath(f.Path()) {
			return true
		}
		val, err := f.ReadSource()
		if err != nil {
			gerr = err
			return false
		}
		return fn(val)
	})
	return gerr
}

// EachActor iterates over each `Actor` data set in the package unless
// the given handler returns false.
func (r *PackReader) EachActor(fn func(*Actor) bool) error {
	var gerr error
	r.EachFile(func(f *ZipFile) bool {
		if !IsActorPath(f.Path()) {
			return true
		}
		val, err := f.ReadActor()
		if err != nil {
			gerr = err
			return false
		}
		return fn(val)
	})
	return gerr
}

// EachUnitGroup iterates over each `UnitGroup` data set in the package unless
// the given handler returns false.
func (r *PackReader) EachUnitGroup(fn func(*UnitGroup) bool) error {
	var gerr error
	r.EachFile(func(f *ZipFile) bool {
		if !IsUnitGroupPath(f.Path()) {
			return true
		}
		val, err := f.ReadUnitGroup()
		if err != nil {
			gerr = err
			return false
		}
		return fn(val)
	})
	return gerr
}

// EachFlowProperty iterates over each `FlowProperty` data set in the package unless
// the given handler returns false.
func (r *PackReader) EachFlowProperty(fn func(*FlowProperty) bool) error {
	var gerr error
	r.EachFile(func(f *ZipFile) bool {
		if !IsFlowPropertyPath(f.Path()) {
			return true
		}
		val, err := f.ReadFlowProperty()
		if err != nil {
			gerr = err
			return false
		}
		return fn(val)
	})
	return gerr
}

// EachFlow iterates over each `Flow` data set in the package unless
// the given handler returns false.
func (r *PackReader) EachFlow(fn func(*Flow) bool) error {
	var gerr error
	r.EachFile(func(f *ZipFile) bool {
		if !IsFlowPath(f.Path()) {
			return true
		}
		val, err := f.ReadFlow()
		if err != nil {
			gerr = err
			return false
		}
		return fn(val)
	})
	return gerr
}

// EachProcess iterates over each `Process` data set in the package unless
// the given handler returns false.
func (r *PackReader) EachProcess(fn func(*Process) bool) error {
	var gerr error
	r.EachFile(func(f *ZipFile) bool {
		if !IsProcessPath(f.Path()) {
			return true
		}
		val, err := f.ReadProcess()
		if err != nil {
			gerr = err
			return false
		}
		return fn(val)
	})
	return gerr
}

// EachImpactCategory iterates over each `ImpactCategory` data set in the package unless
// the given handler returns false.
func (r *PackReader) EachImpactCategory(fn func(*ImpactCategory) bool) error {
	var gerr error
	r.EachFile(func(f *ZipFile) bool {
		if !IsImpactCategoryPath(f.Path()) {
			return true
		}
		val, err := f.ReadImpactCategory()
		if err != nil {
			gerr = err
			return false
		}
		return fn(val)
	})
	return gerr
}

// EachImpactMethod iterates over each `ImpactMethod` data set in the package unless
// the given handler returns false.
func (r *PackReader) EachImpactMethod(fn func(*ImpactMethod) bool) error {
	var gerr error
	r.EachFile(func(f *ZipFile) bool {
		if !IsImpactMethodPath(f.Path()) {
			return true
		}
		val, err := f.ReadImpactMethod()
		if err != nil {
			gerr = err
			return false
		}
		return fn(val)
	})
	return gerr
}

// EachLocation iterates over each `Location` data set in the package unless
// the given handler returns false.
func (r *PackReader) EachLocation(fn func(*Location) bool) error {
	var gerr error
	r.EachFile(func(f *ZipFile) bool {
		if !IsLocationPath(f.Path()) {
			return true
		}
		val, err := f.ReadLocation()
		if err != nil {
			gerr = err
			return false
		}
		return fn(val)
	})
	return gerr
}

// EachParameter iterates over each `Parameter` data set in the package unless
// the given handler returns false.
func (r *PackReader) EachParameter(fn func(*Parameter) bool) error {
	var gerr error
	r.EachFile(func(f *ZipFile) bool {
		if !IsParameterPath(f.Path()) {
			return true
		}
		val, err := f.ReadParameter()
		if err != nil {
			gerr = err
			return false
		}
		return fn(val)
	})
	return gerr
}

// EachSocialIndicator iterates over each `SocialIndicator` data set in the package unless
// the given handler returns false.
func (r *PackReader) EachSocialIndicator(fn func(*SocialIndicator) bool) error {
	var gerr error
	r.EachFile(func(f *ZipFile) bool {
		if !IsSocialIndicatorPath(f.Path()) {
			return true
		}
		val, err := f.ReadSocialIndicator()
		if err != nil {
			gerr = err
			return false
		}
		return fn(val)
	})
	return gerr
}

// EachProductSystem iterates over each `ProductSystem` data set in the package unless
// the given handler returns false.
func (r *PackReader) EachProductSystem(fn func(*ProductSystem) bool) error {
	var gerr error
	r.EachFile(func(f *ZipFile) bool {
		if !IsProductSystemPath(f.Path()) {
			return true
		}
		val, err := f.ReadProductSystem()
		if err != nil {
			gerr = err
			return false
		}
		return fn(val)
	})
	return gerr
}
