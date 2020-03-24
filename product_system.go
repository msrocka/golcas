package golcas

// ProductSystem http://greendelta.github.io/olca-schema/html/ProductSystem.html
type ProductSystem struct {
	CategorizedEntity
	ReferenceProcess *Ref `json:"referenceProcess,omitempty"`

	// The quantitative reference of the product system
	// is an exchange of the reference process.
	ReferenceExchange *ExchangeRef `json:"referenceExchange,omitempty"`

	// The flow property of the reference flow.
	ReferenceFlowProperty *Ref `json:"targetFlowProperty,omitempty"`

	// The unit of the reference flow.
	ReferenceUnit *Ref `json:"targetUnit,omitempty"`

	// The amount of the reference flow.
	ReferenceAmount float64 `json:"targetAmount"`

	// References to the processes of the product system
	Processes []*Ref `json:"processes"`

	ProcessLinks []*ProcessLink `json:"processLinks"`
}

// A ProcessLink describes a link between two processes or between
// a process and a sub-system in a product system.
type ProcessLink struct {
	IsSystemLink bool         `json:"isSystemLink"`
	Provider     *Ref         `json:"provider,omitempty"`
	Flow         *FlowRef     `json:"flow,omitempty"`
	Process      *Ref         `json:"process,omitempty"`
	Exchange     *ExchangeRef `json:"exchange,omitempty"`
}

// An ExchangeRef stores a reference of an exchange of a process.
// Exchanges have an internal ID within a process. We just need
// to store that ID to identify the respective process exchange.
type ExchangeRef struct {
	InternalID int `json:"internalId"`
}
