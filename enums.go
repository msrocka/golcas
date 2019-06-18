package golcas

// ModelType enumeration
type ModelType string

// Enum constants of ModelType
const (
	ProjectModel         ModelType = "PROJECT"
	ImpactMethodModel    ModelType = "IMPACT_METHOD"
	ImpactCategoryModel  ModelType = "IMPACT_CATEGORY"
	ProductSystemModel   ModelType = "PRODUCT_SYSTEM"
	ProcessModel         ModelType = "PROCESS"
	FlowModel            ModelType = "FLOW"
	FlowPropertyModel    ModelType = "FLOW_PROPERTY"
	UnitGroupModel       ModelType = "UNIT_GROUP"
	UnitModel            ModelType = "UNIT"
	ActorModel           ModelType = "ACTOR"
	SourceModel          ModelType = "SOURCE"
	CategoryModel        ModelType = "CATEGORY"
	LocationModel        ModelType = "LOCATION"
	NwSetModel           ModelType = "NW_SET"
	SocialIndicatorModel ModelType = "SOCIAL_INDICATOR"
)

// FlowPropertyType enumeration
type FlowPropertyType string

// Enum constants of FlowPropertyType
const (
	EconomicQuantity FlowPropertyType = "ECONOMIC_QUANTITY"
	PhysicalQuantity FlowPropertyType = "PHYSICAL_QUANTITY"
)

// FlowType enumeration
type FlowType string

// Enum constants of FlowType
const (
	ElementaryFlow FlowType = "ELEMENTARY_FLOW"
	ProductFlow    FlowType = "PRODUCT_FLOW"
	WasteFlow      FlowType = "WASTE_FLOW"
)

// AllocationType enumeration
type AllocationType string

// Enum constants of AllocationType
const (
	PhysicalAllocation AllocationType = "PHYSICAL_ALLOCATION"
	EconomicAllocation AllocationType = "ECONOMIC_ALLOCATION"
	CausalAllocation   AllocationType = "CAUSAL_ALLOCATION"
)

// UncertaintyType enumeration
type UncertaintyType string

// Enum constants of UncertaintyType
const (
	LogNormalDistribution UncertaintyType = "LOG_NORMAL_DISTRIBUTION"
	NormalDistribution    UncertaintyType = "NORMAL_DISTRIBUTION"
	TriangleDistribution  UncertaintyType = "TRIANGLE_DISTRIBUTION"
	UniformDistribution   UncertaintyType = "UNIFORM_DISTRIBUTION"
)

// ParameterScope enumeration
type ParameterScope string

// Enum constants of ParameterScope
const (
	ProcessScope    ParameterScope = "PROCESS_SCOPE"
	LCIAMethodScope ParameterScope = "LCIA_METHOD_SCOPE"
	GlobalScope     ParameterScope = "GLOBAL_SCOPE"
)

// ProcessType enumeration
type ProcessType string

// Enum constants of ProcessType
const (
	LCIResult   ProcessType = "LCI_RESULT"
	UnitProcess ProcessType = "UNIT_PROCESS"
)

// RiskLevel enumeration
type RiskLevel string

// Enum constants of RiskLevel
const (
	HighOpportunity   RiskLevel = "HIGH_OPPORTUNITY"
	MediumOpportunity RiskLevel = "MEDIUM_OPPORTUNITY"
	LowOpportunity    RiskLevel = "LOW_OPPORTUNITY"
	NoRisk            RiskLevel = "NO_RISK"
	VeryLowRisk       RiskLevel = "VERY_LOW_RISK"
	LowRisk           RiskLevel = "LOW_RISK"
	MediumRisk        RiskLevel = "MEDIUM_RISK"
	HighRisk          RiskLevel = "HIGH_RISK"
	VeryHighRisk      RiskLevel = "VERY_HIGH_RISK"
	NoData            RiskLevel = "NO_DATA"
	NotApplicable     RiskLevel = "NOT_APPLICABLE"
)
