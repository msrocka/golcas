package ld

// NewReference creates a RootEntity with the given id, type, and name.
// This function can be used for creating references to root entities.
func NewReference(eType, id, name string) *RootEntity {
	e := RootEntity{}
	e.Type = eType
	e.ID = id
	e.Name = name
	return &e
}

// NewCategory initializes a new category with the given ID and name.
func NewCategory(id, name string, modelType ModelType) *Category {
	c := Category{}
	c.Context = ContextURL
	c.Type = "Category"
	c.ModelType = modelType
	c.ID = id
	c.Name = name
	return &c
}

// NewActor initializes a new Actor with the given id and name
func NewActor(id, name string) *Actor {
	a := Actor{}
	a.Context = ContextURL
	a.Type = "Actor"
	a.ID = id
	a.Name = name
	return &a
}

// NewSource initializes a new Source with the given id and name
func NewSource(id, name string) *Source {
	s := Source{}
	s.Context = ContextURL
	s.Type = "Source"
	s.ID = id
	s.Name = name
	return &s
}

// NewUnit initializes a new Unit with the given id and name
func NewUnit(id, name string) *Unit {
	u := Unit{}
	u.Context = ContextURL
	u.Type = "Unit"
	u.ID = id
	u.Name = name
	return &u
}

// NewUnitGroup initializes a new UnitGroup with the given id and name
func NewUnitGroup(id, name string) *UnitGroup {
	u := UnitGroup{}
	u.Context = ContextURL
	u.Type = "UnitGroup"
	u.ID = id
	u.Name = name
	return &u
}

// NewFlowProperty initializes a new FlowProperty with the given id and name
func NewFlowProperty(id, name string) *FlowProperty {
	f := FlowProperty{}
	f.Context = ContextURL
	f.Type = "FlowProperty"
	f.ID = id
	f.Name = name
	return &f
}

// NewLocation initializes a new Location with the given id and name
func NewLocation(id, name string) *Location {
	l := Location{}
	l.Context = ContextURL
	l.Type = "Location"
	l.ID = id
	l.Name = name
	return &l
}

// NewFlow initializes a new Flow with the given id and name
func NewFlow(id, name string) *Flow {
	f := Flow{}
	f.Context = ContextURL
	f.Type = "Flow"
	f.ID = id
	f.Name = name
	return &f
}

// NewProcess initializes a new Process with the given id and name
func NewProcess(id, name string) *Process {
	p := Process{}
	p.Context = ContextURL
	p.Type = "Process"
	p.ID = id
	p.Name = name
	return &p
}

// NewSocialIndicator initializes a new SocialIndicator with the given id and name
func NewSocialIndicator(id, name string) *SocialIndicator {
	i := SocialIndicator{}
	i.Context = ContextURL
	i.Type = "SocialIndicator"
	i.ID = id
	i.Name = name
	return &i
}

// NewImpactMethod initializes a new ImpactMethod with the given id and name
func NewImpactMethod(id, name string) *ImpactMethod {
	i := ImpactMethod{}
	i.Context = ContextURL
	i.Type = "ImpactMethod"
	i.ID = id
	i.Name = name
	return &i
}

// NewImpactCategory initializes a new ImpactCategory with the given id and name
func NewImpactCategory(id, name string) *ImpactCategory {
	i := ImpactCategory{}
	i.Context = ContextURL
	i.Type = "ImpactCategory"
	i.ID = id
	i.Name = name
	return &i
}
