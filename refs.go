package golcas

// Ref http://greendelta.github.io/olca-schema/html/Ref.html
type Ref struct {
	RootEntity
	CategoryPath []string `json:"categoryPath,omitempty"`
}

// FlowRef http://greendelta.github.io/olca-schema/html/FlowRef.html
type FlowRef struct {
	Ref
	RefUnit  string   `json:"refUnit,omitempty"`
	Location string   `json:"location,omitempty"`
	FlowType FlowType `json:"flowType,omitempty"`
}

// AsRef converts the given flow into a flow reference.
func (f *Flow) AsRef() *FlowRef {
	r := &FlowRef{}
	if f == nil {
		return r
	}
	r.ID = f.ID
	r.Name = f.Name
	if f.Category != nil && f.Category.CategoryPath != nil {
		path := make([]string, len(f.Category.CategoryPath))
		copy(path, f.Category.CategoryPath)
		r.CategoryPath = path
	}
	r.Description = f.Description
	if f.Location != nil {
		r.Location = f.Location.Name
	}
	r.FlowType = f.Type
	r.LastChange = f.LastChange
	r.Version = f.Version
	return r
}
