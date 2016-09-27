package ld

// FlowPropertyFactor http://greendelta.github.io/olca-schema/html/FlowPropertyFactor.html
type FlowPropertyFactor struct {
	Entity
	FlowProperty          *RootEntity `json:"flowProperty,omitempty"`
	ConversionFactor      float64     `json:"conversionFactor"`
	ReferenceFlowProperty bool        `json:"referenceFlowProperty"`
}

// Flow http://greendelta.github.io/olca-schema/html/Flow.html
type Flow struct {
	CategorizedEntity
	Type           FlowType             `json:"flowType,omitempty"`
	Cas            string               `json:"cas,omitempty"`
	Formula        string               `json:"formula,omitempty"`
	FlowProperties []FlowPropertyFactor `json:"flowProperties,omitempty"`
	Location       *RootEntity          `json:"location,omitempty"`
}
