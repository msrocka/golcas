// Package schema provides a Go implementation of the olca-schema definition
// (http://greendelta.github.io/olca-schema).
package schema

// ContextURL is the URL to the olca-schema definition
const ContextURL = "http://greendelta.github.io/olca-schema/context.jsonld"

// Entity http://greendelta.github.io/olca-schema/html/Entity.html
type Entity struct {
	Context string `json:"@context,omitempty"`
	ID      string `json:"@id,omitempty"`
	Type    string `json:"@type,omitempty"`
}

// RootEntity http://greendelta.github.io/olca-schema/html/RootEntity.html
type RootEntity struct {
	Entity
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Version     string `json:"version,omitempty"`
	LastChange  string `json:"lastChange,omitempty"`
}

// CategorizedEntity http://greendelta.github.io/olca-schema/html/CategorizedEntity.html
type CategorizedEntity struct {
	RootEntity
	Category *RootEntity `json:"category,omitempty"`
}

// Category http://greendelta.github.io/olca-schema/html/Category.html
type Category struct {
	CategorizedEntity
	ModelType ModelType `json:"modelType,omitempty"`
}

// Actor http://greendelta.github.io/olca-schema/html/Actor.html
type Actor struct {
	CategorizedEntity
	Address   string `json:"address,omitempty"`
	City      string `json:"city,omitempty"`
	Country   string `json:"country,omitempty"`
	Email     string `json:"email,omitempty"`
	Telefax   string `json:"telefax,omitempty"`
	Telephone string `json:"telephone,omitempty"`
	Website   string `json:"website,omitempty"`
	ZipCode   string `json:"zipCode,omitempty"`
}

// Source http://greendelta.github.io/olca-schema/html/Source.html
type Source struct {
	CategorizedEntity
	Doi           string `json:"doi,omitempty"`
	TextReference string `json:"textReference,omitempty"`
	Year          int    `json:"year,omitempty"`
	ExternalFile  string `json:"externalFile,omitempty"`
}

// Unit http://greendelta.github.io/olca-schema/html/Unit.html
type Unit struct {
	RootEntity
	ConversionFactor float64  `json:"conversionFactor"`
	ReferenceUnit    bool     `json:"referenceUnit"`
	Synonyms         []string `json:"synonyms,omitempty"`
}

// UnitGroup http://greendelta.github.io/olca-schema/html/UnitGroup.html
type UnitGroup struct {
	CategorizedEntity
	DefaultFlowProperty *RootEntity `json:"defaultFlowProperty,omitempty"`
	Units               []Unit      `json:"units,omitempty"`
}

// FlowProperty http://greendelta.github.io/olca-schema/html/FlowProperty.html
type FlowProperty struct {
	CategorizedEntity
	Type      FlowPropertyType `json:"flowPropertyType,omitempty"`
	UnitGroup *RootEntity      `json:"unitGroup,omitempty"`
}

// Location http://greendelta.github.io/olca-schema/html/Location.html
type Location struct {
	RootEntity
	Code      string  `json:"code,omitempty"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Kml       string  `json:"kml,omitempty"`
}

// Uncertainty http://greendelta.github.io/olca-schema/html/Uncertainty.html
type Uncertainty struct {
	Entity
	Type            UncertaintyType `json:"distributionType,omitempty"`
	Mean            *float64        `json:"mean,omitempty"`
	MeanFormula     string          `json:"meanFormula,omitempty"`
	GeomMean        *float64        `json:"geomMean,omitempty"`
	GeomMeanFormula string          `json:"geomMeanFormula,omitempty"`
	Minimum         *float64        `json:"minimum,omitempty"`
	MinimumFormula  string          `json:"minimumFormula,omitempty"`
	Sd              *float64        `json:"sd,omitempty"`
	SdFormula       string          `json:"sdFormula,omitempty"`
	GeomSd          *float64        `json:"geomSd,omitempty"`
	GeomSdFormula   string          `json:"geomSdFormula,omitempty"`
	Mode            *float64        `json:"mode,omitempty"`
	ModeFormula     string          `json:"modeFormula,omitempty"`
	Maximum         *float64        `json:"maximum,omitempty"`
	MaximumFormula  string          `json:"maximumFormula,omitempty"`
}

// Parameter http://greendelta.github.io/olca-schema/html/Parameter.html
type Parameter struct {
	Entity
	Name           string         `json:"name,omitempty"`
	Description    string         `json:"description,omitempty"`
	Scope          ParameterScope `json:"parameterScope,omitempty"`
	InputParameter bool           `json:"inputParameter"`
	Value          float64        `json:"value"`
	Formula        string         `json:"formula,omitempty"`
	ExternalSource string         `json:"externalSource,omitempty"`
	SourceType     string         `json:"sourceType,omitempty"`
	Uncertainty    *Uncertainty   `json:"uncertainty,omitempty"`
}

// SocialIndicator http://greendelta.github.io/olca-schema/html/SocialIndicator.html
type SocialIndicator struct {
	CategorizedEntity
	ActivityVariable  string      `json:"activityVariable,omitempty"`
	ActivityQuantity  *RootEntity `json:"activityQuantity,omitempty"`
	ActivityUnit      *RootEntity `json:"activityUnit,omitempty"`
	UnitOfMeasurement string      `json:"unitOfMeasurement,omitempty"`
	EvaluationScheme  string      `json:"evaluationScheme,omitempty"`
}
