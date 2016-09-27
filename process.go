package ld

// ProcessDocumentation http://greendelta.github.io/olca-schema/html/ProcessDocumentation.html
type ProcessDocumentation struct {
	Entity
	TimeDescription              string       `json:"timeDescription,omitempty"`
	ValidUntil                   string       `json:"validUntil,omitempty"`
	ValidFrom                    string       `json:"validFrom,omitempty"`
	TechnologyDescription        string       `json:"technologyDescription,omitempty"`
	DataCollectionDescription    string       `json:"dataCollectionDescription,omitempty"`
	CompletenessDescription      string       `json:"completenessDescription,omitempty"`
	DataSelectionDescription     string       `json:"dataSelectionDescription,omitempty"`
	ReviewDetails                string       `json:"reviewDetails,omitempty"`
	DataTreatmentDescription     string       `json:"dataTreatmentDescription,omitempty"`
	InventoryMethodDescription   string       `json:"inventoryMethodDescription,omitempty"`
	ModelingConstantsDescription string       `json:"modelingConstantsDescription,omitempty"`
	Reviewer                     *RootEntity  `json:"reviewer,omitempty"`
	SamplingDescription          string       `json:"samplingDescription,omitempty"`
	Sources                      []RootEntity `json:"sources,omitempty"`
	RestrictionsDescription      string       `json:"restrictionsDescription,omitempty"`
	Copyright                    bool         `json:"copyright"`
	CreationDate                 string       `json:"creationDate,omitempty"`
	DataDocumentor               *RootEntity  `json:"dataDocumentor,omitempty"`
	DataGenerator                *RootEntity  `json:"dataGenerator,omitempty"`
	DataSetOwner                 *RootEntity  `json:"dataSetOwner,omitempty"`
	IntendedApplication          string       `json:"intendedApplication,omitempty"`
	ProjectDescription           string       `json:"projectDescription,omitempty"`
	Publication                  *RootEntity  `json:"publication,omitempty"`
	GeographyDescription         string       `json:"geographyDescription,omitempty"`
}

// Exchange http://greendelta.github.io/olca-schema/html/Exchange.html
type Exchange struct {
	Entity
	AvoidedProduct        bool         `json:"avoidedProduct"`
	Flow                  *RootEntity  `json:"flow,omitempty"`
	FlowProperty          *RootEntity  `json:"flowProperty,omitempty"`
	Input                 bool         `json:"input"`
	QuantitativeReference bool         `json:"quantitativeReference"`
	BaseUncertainty       float64      `json:"baseUncertainty"`
	Provider              *RootEntity  `json:"provider,omitempty"`
	Amount                float64      `json:"amount"`
	AmountFormula         string       `json:"amountFormula,omitempty"`
	Unit                  *RootEntity  `json:"unit,omitempty"`
	PedigreeUncertainty   string       `json:"pedigreeUncertainty,omitempty"`
	Uncertainty           *Uncertainty `json:"uncertainty,omitempty"`
	Comment               string       `json:"comment,omitempty"`
}

// AllocationFactor http://greendelta.github.io/olca-schema/html/AllocationFactor.html
type AllocationFactor struct {
	Entity
	ProductExchange   *Entity        `json:"productExchange,omitempty"`
	Type              AllocationType `json:"allocationType,omitempty"`
	Value             float64        `json:"value"`
	AllocatedExchange *Entity        `json:"allocatedExchange,omitempty"`
}

// SocialAspect http://greendelta.github.io/olca-schema/html/SocialAspect.html
type SocialAspect struct {
	Entity
	Indicator     *RootEntity `json:"socialIndicator,omitempty"`
	ActivityValue float64     `json:"activityValue"`
	RawAmount     string      `json:"rawAmount,omitempty"`
	RiskLevel     RiskLevel   `json:"riskLevel,omitempty"`
	Comment       string      `json:"comment,omitempty"`
	Source        *RootEntity `json:"source,omitempty"`
	Quality       string      `json:"quality,omitempty"`
}

// Process http://greendelta.github.io/olca-schema/html/Process.html
type Process struct {
	CategorizedEntity
	DefaultAllocationMethod AllocationType       `json:"defaultAllocationMethod,omitempty"`
	AllocationFactors       []AllocationFactor   `json:"allocationFactors,omitempty"`
	Exchanges               []Exchange           `json:"exchanges,omitempty"`
	Location                *RootEntity          `json:"location,omitempty"`
	Parameters              []RootEntity         `json:"parameters,omitempty"`
	ProcessDocumentation    ProcessDocumentation `json:"processDocumentation"`
	Type                    ProcessType          `json:"processTyp,omitempty"`
	SocialAspects           []SocialAspect       `json:"socialAspects,omitempty"`
}
