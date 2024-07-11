// DO NOT CHANGE THIS FILE AS THIS WAS GENERATED AUTOMATICALLY
// instead modify the generator that produces this file, see
// http://greendelta.github.io/olca-schema

package golcas

type AllocationType string
const(
	PHYSICAL_ALLOCATION AllocationType = "PHYSICAL_ALLOCATION"
	ECONOMIC_ALLOCATION AllocationType = "ECONOMIC_ALLOCATION"
	CAUSAL_ALLOCATION AllocationType = "CAUSAL_ALLOCATION"
	USE_DEFAULT_ALLOCATION AllocationType = "USE_DEFAULT_ALLOCATION"
	NO_ALLOCATION AllocationType = "NO_ALLOCATION"
)

type Direction string
const(
	INPUT Direction = "INPUT"
	OUTPUT Direction = "OUTPUT"
)

type FlowPropertyType string
const(
	ECONOMIC_QUANTITY FlowPropertyType = "ECONOMIC_QUANTITY"
	PHYSICAL_QUANTITY FlowPropertyType = "PHYSICAL_QUANTITY"
)

type FlowType string
const(
	ELEMENTARY_FLOW FlowType = "ELEMENTARY_FLOW"
	PRODUCT_FLOW FlowType = "PRODUCT_FLOW"
	WASTE_FLOW FlowType = "WASTE_FLOW"
)

type ModelType string
const(
	ACTOR ModelType = "ACTOR"
	CATEGORY ModelType = "CATEGORY"
	CURRENCY ModelType = "CURRENCY"
	DQ_SYSTEM ModelType = "DQ_SYSTEM"
	EPD ModelType = "EPD"
	FLOW ModelType = "FLOW"
	FLOW_PROPERTY ModelType = "FLOW_PROPERTY"
	IMPACT_CATEGORY ModelType = "IMPACT_CATEGORY"
	IMPACT_METHOD ModelType = "IMPACT_METHOD"
	LOCATION ModelType = "LOCATION"
	PARAMETER ModelType = "PARAMETER"
	PROCESS ModelType = "PROCESS"
	PRODUCT_SYSTEM ModelType = "PRODUCT_SYSTEM"
	PROJECT ModelType = "PROJECT"
	RESULT ModelType = "RESULT"
	SOCIAL_INDICATOR ModelType = "SOCIAL_INDICATOR"
	SOURCE ModelType = "SOURCE"
	UNIT_GROUP ModelType = "UNIT_GROUP"
)

type ParameterScope string
const(
	PROCESS_SCOPE ParameterScope = "PROCESS_SCOPE"
	IMPACT_SCOPE ParameterScope = "IMPACT_SCOPE"
	GLOBAL_SCOPE ParameterScope = "GLOBAL_SCOPE"
)

type ProcessType string
const(
	LCI_RESULT ProcessType = "LCI_RESULT"
	UNIT_PROCESS ProcessType = "UNIT_PROCESS"
)

type ProviderLinking string
const(
	IGNORE_DEFAULTS ProviderLinking = "IGNORE_DEFAULTS"
	PREFER_DEFAULTS ProviderLinking = "PREFER_DEFAULTS"
	ONLY_DEFAULTS ProviderLinking = "ONLY_DEFAULTS"
)

type RiskLevel string
const(
	NO_OPPORTUNITY RiskLevel = "NO_OPPORTUNITY"
	HIGH_OPPORTUNITY RiskLevel = "HIGH_OPPORTUNITY"
	MEDIUM_OPPORTUNITY RiskLevel = "MEDIUM_OPPORTUNITY"
	LOW_OPPORTUNITY RiskLevel = "LOW_OPPORTUNITY"
	NO_RISK RiskLevel = "NO_RISK"
	VERY_LOW_RISK RiskLevel = "VERY_LOW_RISK"
	LOW_RISK RiskLevel = "LOW_RISK"
	MEDIUM_RISK RiskLevel = "MEDIUM_RISK"
	HIGH_RISK RiskLevel = "HIGH_RISK"
	VERY_HIGH_RISK RiskLevel = "VERY_HIGH_RISK"
	NO_DATA RiskLevel = "NO_DATA"
	NOT_APPLICABLE RiskLevel = "NOT_APPLICABLE"
)

type UncertaintyType string
const(
	LOG_NORMAL_DISTRIBUTION UncertaintyType = "LOG_NORMAL_DISTRIBUTION"
	NORMAL_DISTRIBUTION UncertaintyType = "NORMAL_DISTRIBUTION"
	TRIANGLE_DISTRIBUTION UncertaintyType = "TRIANGLE_DISTRIBUTION"
	UNIFORM_DISTRIBUTION UncertaintyType = "UNIFORM_DISTRIBUTION"
)

type Actor struct {
	ID string `json:"@id,omitempty"`
	Address string `json:"address,omitempty"`
	Category string `json:"category,omitempty"`
	City string `json:"city,omitempty"`
	Country string `json:"country,omitempty"`
	Description string `json:"description,omitempty"`
	Email string `json:"email,omitempty"`
	LastChange string `json:"lastChange"`
	Name string `json:"name,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Telefax string `json:"telefax,omitempty"`
	Telephone string `json:"telephone,omitempty"`
	Version string `json:"version,omitempty"`
	Website string `json:"website,omitempty"`
	ZipCode string `json:"zipCode,omitempty"`
}

type AllocationFactor struct {
	AllocationType AllocationType `json:"allocationType,omitempty"`
	Exchange *ExchangeRef `json:"exchange"`
	Formula string `json:"formula,omitempty"`
	Product *Ref `json:"product"`
	Value float64 `json:"value"`
}

type Currency struct {
	ID string `json:"@id,omitempty"`
	Category string `json:"category,omitempty"`
	Code string `json:"code,omitempty"`
	ConversionFactor float64 `json:"conversionFactor"`
	Description string `json:"description,omitempty"`
	LastChange string `json:"lastChange"`
	Name string `json:"name,omitempty"`
	RefCurrency *Ref `json:"refCurrency"`
	Tags []string `json:"tags,omitempty"`
	Version string `json:"version,omitempty"`
}

type DQIndicator struct {
	Name string `json:"name,omitempty"`
	Position int `json:"position"`
	Scores []DQScore `json:"scores,omitempty"`
}

type DQScore struct {
	Description string `json:"description,omitempty"`
	Label string `json:"label,omitempty"`
	Position int `json:"position"`
	Uncertainty *float64 `json:"uncertainty"`
}

type DQSystem struct {
	ID string `json:"@id,omitempty"`
	Category string `json:"category,omitempty"`
	Description string `json:"description,omitempty"`
	HasUncertainties bool `json:"hasUncertainties"`
	Indicators []DQIndicator `json:"indicators,omitempty"`
	LastChange string `json:"lastChange"`
	Name string `json:"name,omitempty"`
	Source *Ref `json:"source"`
	Tags []string `json:"tags,omitempty"`
	Version string `json:"version,omitempty"`
}

type Epd struct {
	ID string `json:"@id,omitempty"`
	Category string `json:"category,omitempty"`
	Description string `json:"description,omitempty"`
	LastChange string `json:"lastChange"`
	Manufacturer *Ref `json:"manufacturer"`
	Modules []EpdModule `json:"modules,omitempty"`
	Name string `json:"name,omitempty"`
	Pcr *Ref `json:"pcr"`
	Product *EpdProduct `json:"product"`
	ProgramOperator *Ref `json:"programOperator"`
	Tags []string `json:"tags,omitempty"`
	Urn string `json:"urn,omitempty"`
	Verifier *Ref `json:"verifier"`
	Version string `json:"version,omitempty"`
}

type EpdModule struct {
	Multiplier float64 `json:"multiplier"`
	Name string `json:"name,omitempty"`
	Result *Ref `json:"result"`
}

type EpdProduct struct {
	Amount float64 `json:"amount"`
	Flow *Ref `json:"flow"`
	FlowProperty *Ref `json:"flowProperty"`
	Unit *Ref `json:"unit"`
}

type Exchange struct {
	Amount float64 `json:"amount"`
	AmountFormula string `json:"amountFormula,omitempty"`
	BaseUncertainty *float64 `json:"baseUncertainty"`
	CostFormula string `json:"costFormula,omitempty"`
	CostValue *float64 `json:"costValue"`
	Currency *Ref `json:"currency"`
	DefaultProvider *Ref `json:"defaultProvider"`
	Description string `json:"description,omitempty"`
	DqEntry string `json:"dqEntry,omitempty"`
	Flow *Ref `json:"flow"`
	FlowProperty *Ref `json:"flowProperty"`
	InternalId int `json:"internalId"`
	IsAvoidedProduct bool `json:"isAvoidedProduct"`
	IsInput bool `json:"isInput"`
	IsQuantitativeReference bool `json:"isQuantitativeReference"`
	Location *Ref `json:"location"`
	Uncertainty *Uncertainty `json:"uncertainty"`
	Unit *Ref `json:"unit"`
}

type ExchangeRef struct {
	InternalId int `json:"internalId"`
}

type Flow struct {
	ID string `json:"@id,omitempty"`
	Cas string `json:"cas,omitempty"`
	Category string `json:"category,omitempty"`
	Description string `json:"description,omitempty"`
	FlowProperties []FlowPropertyFactor `json:"flowProperties,omitempty"`
	FlowType FlowType `json:"flowType,omitempty"`
	Formula string `json:"formula,omitempty"`
	IsInfrastructureFlow bool `json:"isInfrastructureFlow"`
	LastChange string `json:"lastChange"`
	Location *Ref `json:"location"`
	Name string `json:"name,omitempty"`
	Synonyms string `json:"synonyms,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Version string `json:"version,omitempty"`
}

type FlowMap struct {
	ID string `json:"@id,omitempty"`
	Category string `json:"category,omitempty"`
	Description string `json:"description,omitempty"`
	LastChange string `json:"lastChange"`
	Mappings []FlowMapEntry `json:"mappings,omitempty"`
	Name string `json:"name,omitempty"`
	Source *Ref `json:"source"`
	Tags []string `json:"tags,omitempty"`
	Target *Ref `json:"target"`
	Version string `json:"version,omitempty"`
}

type FlowMapEntry struct {
	ConversionFactor float64 `json:"conversionFactor"`
	From *FlowMapRef `json:"from"`
	To *FlowMapRef `json:"to"`
}

type FlowMapRef struct {
	Flow *Ref `json:"flow"`
	FlowProperty *Ref `json:"flowProperty"`
	Provider *Ref `json:"provider"`
	Unit *Ref `json:"unit"`
}

type FlowProperty struct {
	ID string `json:"@id,omitempty"`
	Category string `json:"category,omitempty"`
	Description string `json:"description,omitempty"`
	FlowPropertyType FlowPropertyType `json:"flowPropertyType,omitempty"`
	LastChange string `json:"lastChange"`
	Name string `json:"name,omitempty"`
	Tags []string `json:"tags,omitempty"`
	UnitGroup *Ref `json:"unitGroup"`
	Version string `json:"version,omitempty"`
}

type FlowPropertyFactor struct {
	ConversionFactor float64 `json:"conversionFactor"`
	FlowProperty *Ref `json:"flowProperty"`
	IsRefFlowProperty bool `json:"isRefFlowProperty"`
}

type FlowResult struct {
	Amount float64 `json:"amount"`
	Description string `json:"description,omitempty"`
	Flow *Ref `json:"flow"`
	FlowProperty *Ref `json:"flowProperty"`
	IsInput bool `json:"isInput"`
	IsRefFlow bool `json:"isRefFlow"`
	Location *Ref `json:"location"`
	Unit *Ref `json:"unit"`
}

type ImpactCategory struct {
	ID string `json:"@id,omitempty"`
	Category string `json:"category,omitempty"`
	Code string `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
	Direction Direction `json:"direction,omitempty"`
	ImpactFactors []ImpactFactor `json:"impactFactors,omitempty"`
	LastChange string `json:"lastChange"`
	Name string `json:"name,omitempty"`
	Parameters []Parameter `json:"parameters,omitempty"`
	RefUnit string `json:"refUnit,omitempty"`
	Source *Ref `json:"source"`
	Tags []string `json:"tags,omitempty"`
	Version string `json:"version,omitempty"`
}

type ImpactFactor struct {
	Flow *Ref `json:"flow"`
	FlowProperty *Ref `json:"flowProperty"`
	Formula string `json:"formula,omitempty"`
	Location *Ref `json:"location"`
	Uncertainty *Uncertainty `json:"uncertainty"`
	Unit *Ref `json:"unit"`
	Value float64 `json:"value"`
}

type ImpactMethod struct {
	ID string `json:"@id,omitempty"`
	Category string `json:"category,omitempty"`
	Code string `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
	ImpactCategories []Ref `json:"impactCategories,omitempty"`
	LastChange string `json:"lastChange"`
	Name string `json:"name,omitempty"`
	NwSets []NwSet `json:"nwSets,omitempty"`
	Source *Ref `json:"source"`
	Tags []string `json:"tags,omitempty"`
	Version string `json:"version,omitempty"`
}

type ImpactResult struct {
	Amount float64 `json:"amount"`
	Description string `json:"description,omitempty"`
	Indicator *Ref `json:"indicator"`
}

type LinkingConfig struct {
	Cutoff *float64 `json:"cutoff"`
	PreferUnitProcesses bool `json:"preferUnitProcesses"`
	ProviderLinking ProviderLinking `json:"providerLinking,omitempty"`
}

type Location struct {
	ID string `json:"@id,omitempty"`
	Category string `json:"category,omitempty"`
	Code string `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
	Geometry map[string]any `json:"geometry"`
	LastChange string `json:"lastChange"`
	Latitude *float64 `json:"latitude"`
	Longitude *float64 `json:"longitude"`
	Name string `json:"name,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Version string `json:"version,omitempty"`
}

type NwFactor struct {
	ImpactCategory *Ref `json:"impactCategory"`
	NormalisationFactor *float64 `json:"normalisationFactor"`
	WeightingFactor *float64 `json:"weightingFactor"`
}

type NwSet struct {
	ID string `json:"@id,omitempty"`
	Description string `json:"description,omitempty"`
	Factors []NwFactor `json:"factors,omitempty"`
	Name string `json:"name,omitempty"`
	WeightedScoreUnit string `json:"weightedScoreUnit,omitempty"`
}

type Parameter struct {
	ID string `json:"@id,omitempty"`
	Category string `json:"category,omitempty"`
	Description string `json:"description,omitempty"`
	Formula string `json:"formula,omitempty"`
	IsInputParameter bool `json:"isInputParameter"`
	LastChange string `json:"lastChange"`
	Name string `json:"name,omitempty"`
	ParameterScope ParameterScope `json:"parameterScope,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Uncertainty *Uncertainty `json:"uncertainty"`
	Value float64 `json:"value"`
	Version string `json:"version,omitempty"`
}

type ParameterRedef struct {
	Context *Ref `json:"context"`
	Description string `json:"description,omitempty"`
	IsProtected bool `json:"isProtected"`
	Name string `json:"name,omitempty"`
	Uncertainty *Uncertainty `json:"uncertainty"`
	Value float64 `json:"value"`
}

type ParameterRedefSet struct {
	Description string `json:"description,omitempty"`
	IsBaseline bool `json:"isBaseline"`
	Name string `json:"name,omitempty"`
	Parameters []ParameterRedef `json:"parameters,omitempty"`
}

type Process struct {
	ID string `json:"@id,omitempty"`
	AllocationFactors []AllocationFactor `json:"allocationFactors,omitempty"`
	Category string `json:"category,omitempty"`
	DefaultAllocationMethod AllocationType `json:"defaultAllocationMethod,omitempty"`
	Description string `json:"description,omitempty"`
	DqEntry string `json:"dqEntry,omitempty"`
	DqSystem *Ref `json:"dqSystem"`
	ExchangeDqSystem *Ref `json:"exchangeDqSystem"`
	Exchanges []Exchange `json:"exchanges,omitempty"`
	IsInfrastructureProcess bool `json:"isInfrastructureProcess"`
	LastChange string `json:"lastChange"`
	LastInternalId int `json:"lastInternalId"`
	Location *Ref `json:"location"`
	Name string `json:"name,omitempty"`
	Parameters []Parameter `json:"parameters,omitempty"`
	ProcessDocumentation *ProcessDocumentation `json:"processDocumentation"`
	ProcessType ProcessType `json:"processType,omitempty"`
	SocialAspects []SocialAspect `json:"socialAspects,omitempty"`
	SocialDqSystem *Ref `json:"socialDqSystem"`
	Tags []string `json:"tags,omitempty"`
	Version string `json:"version,omitempty"`
}

type ProcessDocumentation struct {
	CompletenessDescription string `json:"completenessDescription,omitempty"`
	CreationDate string `json:"creationDate"`
	DataCollectionDescription string `json:"dataCollectionDescription,omitempty"`
	DataDocumentor *Ref `json:"dataDocumentor"`
	DataGenerator *Ref `json:"dataGenerator"`
	DataSelectionDescription string `json:"dataSelectionDescription,omitempty"`
	DataSetOwner *Ref `json:"dataSetOwner"`
	DataTreatmentDescription string `json:"dataTreatmentDescription,omitempty"`
	GeographyDescription string `json:"geographyDescription,omitempty"`
	IntendedApplication string `json:"intendedApplication,omitempty"`
	InventoryMethodDescription string `json:"inventoryMethodDescription,omitempty"`
	IsCopyrightProtected bool `json:"isCopyrightProtected"`
	ModelingConstantsDescription string `json:"modelingConstantsDescription,omitempty"`
	ProjectDescription string `json:"projectDescription,omitempty"`
	Publication *Ref `json:"publication"`
	RestrictionsDescription string `json:"restrictionsDescription,omitempty"`
	ReviewDetails string `json:"reviewDetails,omitempty"`
	Reviewer *Ref `json:"reviewer"`
	SamplingDescription string `json:"samplingDescription,omitempty"`
	Sources []Ref `json:"sources,omitempty"`
	TechnologyDescription string `json:"technologyDescription,omitempty"`
	TimeDescription string `json:"timeDescription,omitempty"`
	UseAdvice string `json:"useAdvice,omitempty"`
	ValidFrom string `json:"validFrom"`
	ValidUntil string `json:"validUntil"`
}

type ProcessLink struct {
	Exchange *ExchangeRef `json:"exchange"`
	Flow *Ref `json:"flow"`
	Process *Ref `json:"process"`
	Provider *Ref `json:"provider"`
}

type ProductSystem struct {
	ID string `json:"@id,omitempty"`
	Category string `json:"category,omitempty"`
	Description string `json:"description,omitempty"`
	LastChange string `json:"lastChange"`
	Name string `json:"name,omitempty"`
	ParameterSets []ParameterRedefSet `json:"parameterSets,omitempty"`
	ProcessLinks []ProcessLink `json:"processLinks,omitempty"`
	Processes []Ref `json:"processes,omitempty"`
	RefExchange *ExchangeRef `json:"refExchange"`
	RefProcess *Ref `json:"refProcess"`
	Tags []string `json:"tags,omitempty"`
	TargetAmount float64 `json:"targetAmount"`
	TargetFlowProperty *Ref `json:"targetFlowProperty"`
	TargetUnit *Ref `json:"targetUnit"`
	Version string `json:"version,omitempty"`
}

type Project struct {
	ID string `json:"@id,omitempty"`
	Category string `json:"category,omitempty"`
	Description string `json:"description,omitempty"`
	ImpactMethod *Ref `json:"impactMethod"`
	IsWithCosts bool `json:"isWithCosts"`
	IsWithRegionalization bool `json:"isWithRegionalization"`
	LastChange string `json:"lastChange"`
	Name string `json:"name,omitempty"`
	NwSet *NwSet `json:"nwSet"`
	Tags []string `json:"tags,omitempty"`
	Variants []ProjectVariant `json:"variants,omitempty"`
	Version string `json:"version,omitempty"`
}

type ProjectVariant struct {
	AllocationMethod AllocationType `json:"allocationMethod,omitempty"`
	Amount float64 `json:"amount"`
	Description string `json:"description,omitempty"`
	IsDisabled bool `json:"isDisabled"`
	Name string `json:"name,omitempty"`
	ParameterRedefs []ParameterRedef `json:"parameterRedefs,omitempty"`
	ProductSystem *Ref `json:"productSystem"`
	Unit *Ref `json:"unit"`
}

type Ref struct {
	Type string `json:"@type,omitempty"`
	ID string `json:"@id,omitempty"`
	Category string `json:"category,omitempty"`
	Description string `json:"description,omitempty"`
	FlowType FlowType `json:"flowType,omitempty"`
	Location string `json:"location,omitempty"`
	Name string `json:"name,omitempty"`
	ProcessType ProcessType `json:"processType,omitempty"`
	RefUnit string `json:"refUnit,omitempty"`
}

type Result struct {
	ID string `json:"@id,omitempty"`
	Category string `json:"category,omitempty"`
	Description string `json:"description,omitempty"`
	FlowResults []FlowResult `json:"flowResults,omitempty"`
	ImpactMethod *Ref `json:"impactMethod"`
	ImpactResults []ImpactResult `json:"impactResults,omitempty"`
	LastChange string `json:"lastChange"`
	Name string `json:"name,omitempty"`
	ProductSystem *Ref `json:"productSystem"`
	Tags []string `json:"tags,omitempty"`
	Version string `json:"version,omitempty"`
}

type SocialAspect struct {
	ActivityValue float64 `json:"activityValue"`
	Comment string `json:"comment,omitempty"`
	Quality string `json:"quality,omitempty"`
	RawAmount string `json:"rawAmount,omitempty"`
	RiskLevel RiskLevel `json:"riskLevel,omitempty"`
	SocialIndicator *Ref `json:"socialIndicator"`
	Source *Ref `json:"source"`
}

type SocialIndicator struct {
	ID string `json:"@id,omitempty"`
	ActivityQuantity *Ref `json:"activityQuantity"`
	ActivityUnit *Ref `json:"activityUnit"`
	ActivityVariable string `json:"activityVariable,omitempty"`
	Category string `json:"category,omitempty"`
	Description string `json:"description,omitempty"`
	EvaluationScheme string `json:"evaluationScheme,omitempty"`
	LastChange string `json:"lastChange"`
	Name string `json:"name,omitempty"`
	Tags []string `json:"tags,omitempty"`
	UnitOfMeasurement string `json:"unitOfMeasurement,omitempty"`
	Version string `json:"version,omitempty"`
}

type Source struct {
	ID string `json:"@id,omitempty"`
	Category string `json:"category,omitempty"`
	Description string `json:"description,omitempty"`
	ExternalFile string `json:"externalFile,omitempty"`
	LastChange string `json:"lastChange"`
	Name string `json:"name,omitempty"`
	Tags []string `json:"tags,omitempty"`
	TextReference string `json:"textReference,omitempty"`
	Url string `json:"url,omitempty"`
	Version string `json:"version,omitempty"`
	Year *int `json:"year"`
}

type Uncertainty struct {
	DistributionType UncertaintyType `json:"distributionType,omitempty"`
	GeomMean *float64 `json:"geomMean"`
	GeomSd *float64 `json:"geomSd"`
	Maximum *float64 `json:"maximum"`
	Mean *float64 `json:"mean"`
	Minimum *float64 `json:"minimum"`
	Mode *float64 `json:"mode"`
	Sd *float64 `json:"sd"`
}

type Unit struct {
	ID string `json:"@id,omitempty"`
	ConversionFactor float64 `json:"conversionFactor"`
	Description string `json:"description,omitempty"`
	IsRefUnit bool `json:"isRefUnit"`
	Name string `json:"name,omitempty"`
	Synonyms []string `json:"synonyms,omitempty"`
}

type UnitGroup struct {
	ID string `json:"@id,omitempty"`
	Category string `json:"category,omitempty"`
	DefaultFlowProperty *Ref `json:"defaultFlowProperty"`
	Description string `json:"description,omitempty"`
	LastChange string `json:"lastChange"`
	Name string `json:"name,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Units []Unit `json:"units,omitempty"`
	Version string `json:"version,omitempty"`
}

type CalculationSetup struct {
	Allocation AllocationType `json:"allocation,omitempty"`
	Amount float64 `json:"amount"`
	FlowProperty *Ref `json:"flowProperty"`
	ImpactMethod *Ref `json:"impactMethod"`
	NwSet *Ref `json:"nwSet"`
	Parameters []ParameterRedef `json:"parameters,omitempty"`
	Target *Ref `json:"target"`
	Unit *Ref `json:"unit"`
	WithCosts bool `json:"withCosts"`
	WithRegionalization bool `json:"withRegionalization"`
}

type CostValue struct {
	Amount float64 `json:"amount"`
	Currency *Ref `json:"currency"`
}

type EnviFlow struct {
	Flow *Ref `json:"flow"`
	IsInput bool `json:"isInput"`
	Location *Ref `json:"location"`
}

type EnviFlowValue struct {
	Amount float64 `json:"amount"`
	EnviFlow *EnviFlow `json:"enviFlow"`
}

type ImpactValue struct {
	Amount float64 `json:"amount"`
	ImpactCategory *Ref `json:"impactCategory"`
}

type ResultState struct {
	ID string `json:"@id,omitempty"`
	Error string `json:"error,omitempty"`
	IsReady bool `json:"isReady"`
	IsScheduled bool `json:"isScheduled"`
	Time int `json:"time"`
}

type SankeyEdge struct {
	NodeIndex int `json:"nodeIndex"`
	ProviderIndex int `json:"providerIndex"`
	UpstreamShare float64 `json:"upstreamShare"`
}

type SankeyGraph struct {
	Edges []SankeyEdge `json:"edges,omitempty"`
	Nodes []SankeyNode `json:"nodes,omitempty"`
	RootIndex int `json:"rootIndex"`
}

type SankeyNode struct {
	DirectResult float64 `json:"directResult"`
	Index int `json:"index"`
	TechFlow *TechFlow `json:"techFlow"`
	TotalResult float64 `json:"totalResult"`
}

type SankeyRequest struct {
	EnviFlow *EnviFlow `json:"enviFlow"`
	ForCosts *bool `json:"forCosts"`
	ImpactCategory *Ref `json:"impactCategory"`
	MaxNodes *int `json:"maxNodes"`
	MinShare *float64 `json:"minShare"`
}

type TechFlow struct {
	Flow *Ref `json:"flow"`
	Provider *Ref `json:"provider"`
}

type TechFlowValue struct {
	Amount float64 `json:"amount"`
	TechFlow *TechFlow `json:"techFlow"`
}

type UpstreamNode struct {
	DirectContribution float64 `json:"directContribution"`
	RequiredAmount float64 `json:"requiredAmount"`
	Result float64 `json:"result"`
	TechFlow *TechFlow `json:"techFlow"`
}

