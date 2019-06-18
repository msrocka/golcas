package golcas

// ImpactCategory http://greendelta.github.io/olca-schema/html/ImpactCategory.html
type ImpactCategory struct {
	RootEntity
	ReferenceUnit string         `json:"referenceUnitName,omitempty"`
	ImpactFactors []ImpactFactor `json:"impactFactors,omitempty"`
}

// ImpactFactor http://greendelta.github.io/olca-schema/html/ImpactFactor.html
type ImpactFactor struct {
	Entity
	Flow         *FlowRef     `json:"flow,omitempty"`
	FlowProperty *Ref         `json:"flowProperty,omitempty"`
	Unit         *Ref         `json:"unit,omitempty"`
	Value        float64      `json:"value"`
	Formula      string       `json:"formula,omitempty"`
	Uncertainty  *Uncertainty `json:"uncertainty,omitempty"`
}

// ImpactMethod http://greendelta.github.io/olca-schema/html/ImpactMethod.html
type ImpactMethod struct {
	CategorizedEntity
	ImpactCategories []Ref `json:"impactCategories,omitempty"`
	Parameters       []Ref `json:"parameters,omitempty"`
}
