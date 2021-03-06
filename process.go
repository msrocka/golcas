package golcas

// ProcessDocumentation http://greendelta.github.io/olca-schema/html/ProcessDocumentation.html
type ProcessDocumentation struct {
	Entity
	TimeDescription              string `json:"timeDescription,omitempty"`
	ValidUntil                   string `json:"validUntil,omitempty"`
	ValidFrom                    string `json:"validFrom,omitempty"`
	TechnologyDescription        string `json:"technologyDescription,omitempty"`
	DataCollectionDescription    string `json:"dataCollectionDescription,omitempty"`
	CompletenessDescription      string `json:"completenessDescription,omitempty"`
	DataSelectionDescription     string `json:"dataSelectionDescription,omitempty"`
	ReviewDetails                string `json:"reviewDetails,omitempty"`
	DataTreatmentDescription     string `json:"dataTreatmentDescription,omitempty"`
	InventoryMethodDescription   string `json:"inventoryMethodDescription,omitempty"`
	ModelingConstantsDescription string `json:"modelingConstantsDescription,omitempty"`
	Reviewer                     *Ref   `json:"reviewer,omitempty"`
	SamplingDescription          string `json:"samplingDescription,omitempty"`
	Sources                      []Ref  `json:"sources,omitempty"`
	RestrictionsDescription      string `json:"restrictionsDescription,omitempty"`
	HasCopyright                 bool   `json:"copyright"`
	CreationDate                 string `json:"creationDate,omitempty"`
	DataDocumentor               *Ref   `json:"dataDocumentor,omitempty"`
	DataGenerator                *Ref   `json:"dataGenerator,omitempty"`
	DataSetOwner                 *Ref   `json:"dataSetOwner,omitempty"`
	IntendedApplication          string `json:"intendedApplication,omitempty"`
	ProjectDescription           string `json:"projectDescription,omitempty"`
	Publication                  *Ref   `json:"publication,omitempty"`
	GeographyDescription         string `json:"geographyDescription,omitempty"`
}

// Exchange http://greendelta.github.io/olca-schema/html/Exchange.html
type Exchange struct {
	Entity
	InternalID              int          `json:"internalId"`
	IsAvoidedProduct        bool         `json:"avoidedProduct"`
	Flow                    *FlowRef     `json:"flow,omitempty"`
	FlowProperty            *Ref         `json:"flowProperty,omitempty"`
	IsInput                 bool         `json:"input"`
	IsQuantitativeReference bool         `json:"quantitativeReference"`
	BaseUncertainty         float64      `json:"baseUncertainty"`
	Provider                *Ref         `json:"provider,omitempty"` // TODO: ProcessRef
	Amount                  float64      `json:"amount"`
	AmountFormula           string       `json:"amountFormula,omitempty"`
	Unit                    *Ref         `json:"unit,omitempty"`
	PedigreeUncertainty     string       `json:"pedigreeUncertainty,omitempty"`
	Uncertainty             *Uncertainty `json:"uncertainty,omitempty"`
	Comment                 string       `json:"comment,omitempty"`
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
	Indicator     *Ref      `json:"socialIndicator,omitempty"`
	ActivityValue float64   `json:"activityValue"`
	RawAmount     string    `json:"rawAmount,omitempty"`
	RiskLevel     RiskLevel `json:"riskLevel,omitempty"`
	Comment       string    `json:"comment,omitempty"`
	Source        *Ref      `json:"source,omitempty"`
	Quality       string    `json:"quality,omitempty"`
}

// Process http://greendelta.github.io/olca-schema/html/Process.html
type Process struct {
	CategorizedEntity
	DefaultAllocationMethod AllocationType       `json:"defaultAllocationMethod,omitempty"`
	AllocationFactors       []AllocationFactor   `json:"allocationFactors,omitempty"`
	Exchanges               []Exchange           `json:"exchanges,omitempty"`
	Location                *Ref                 `json:"location,omitempty"`
	Parameters              []Ref                `json:"parameters,omitempty"`
	ProcessDocumentation    ProcessDocumentation `json:"processDocumentation"`
	Type                    ProcessType          `json:"processType,omitempty"`
	SocialAspects           []SocialAspect       `json:"socialAspects,omitempty"`
}

// QuantitativeReference reteruns the Exchange that is the
// quantitative reference of the process.
func (p *Process) QuantitativeReference() *Exchange {
	if p == nil {
		return nil
	}
	for i := range p.Exchanges {
		if p.Exchanges[i].IsQuantitativeReference {
			return &p.Exchanges[i]
		}
	}
	return nil
}
