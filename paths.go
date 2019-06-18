package golcas

import (
	"regexp"
	"strings"
)

var uuidRegex *regexp.Regexp

// FindUUID returns the UUID from the given path or an empty string if it cannot
// find it.
func FindUUID(path string) string {
	if uuidRegex == nil {
		pattern := "[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}"
		uuidRegex = regexp.MustCompile(pattern)
	}
	return uuidRegex.FindString(path)
}

// IsCategoryPath returns true if the given path describes a JSON file in the
// folder where data sets of type `Category` are stored.
func IsCategoryPath(path string) bool {
	return isJSONInFolder(path, "categories")
}

// IsSourcePath returns true if the given path describes a JSON file in the
// folder where data sets of type `Source` are stored.
func IsSourcePath(path string) bool {
	return isJSONInFolder(path, "sources")
}

// IsActorPath returns true if the given path describes a JSON file in the
// folder where data sets of type `Actor` are stored.
func IsActorPath(path string) bool {
	return isJSONInFolder(path, "actors")
}

// IsUnitGroupPath returns true if the given path describes a JSON file in the
// folder where data sets of type `UnitGroup` are stored.
func IsUnitGroupPath(path string) bool {
	return isJSONInFolder(path, "unit_groups")
}

// IsFlowPropertyPath returns true if the given path describes a JSON file in the
// folder where data sets of type `FlowProperty` are stored.
func IsFlowPropertyPath(path string) bool {
	return isJSONInFolder(path, "flow_properties")
}

// IsFlowPath returns true if the given path describes a JSON file in the
// folder where data sets of type `Flow` are stored.
func IsFlowPath(path string) bool {
	return isJSONInFolder(path, "flows")
}

// IsProcessPath returns true if the given path describes a JSON file in the
// folder where data sets of type `Process` are stored.
func IsProcessPath(path string) bool {
	return isJSONInFolder(path, "processes")
}

// IsProductSystemPath returns true if the given path describes a JSON file in the
// folder where data sets of type `ProductSystem` are stored.
func IsProductSystemPath(path string) bool {
	return isJSONInFolder(path, "product_systems")
}

// IsCurrencyPath returns true if the given path describes a JSON file in the
// folder where data sets of type `Currency` are stored.
func IsCurrencyPath(path string) bool {
	return isJSONInFolder(path, "currencies")
}

// IsImpactCategoryPath returns true if the given path describes a JSON file in the
// folder where data sets of type `ImpactCategory` are stored.
func IsImpactCategoryPath(path string) bool {
	return isJSONInFolder(path, "impact_categories")
}

// IsImpactMethodPath returns true if the given path describes a JSON file in the
// folder where data sets of type `ImpactMethod` are stored.
func IsImpactMethodPath(path string) bool {
	return isJSONInFolder(path, "impact_methods")
}

// IsLocationPath returns true if the given path describes a JSON file in the
// folder where data sets of type `Location` are stored.
func IsLocationPath(path string) bool {
	return isJSONInFolder(path, "locations")
}

// IsNwSetPath returns true if the given path describes a JSON file in the
// folder where data sets of type `NwSet` are stored.
func IsNwSetPath(path string) bool {
	return isJSONInFolder(path, "nw_sets")
}

// IsParameterPath returns true if the given path describes a JSON file in the
// folder where data sets of type `Parameter` are stored.
func IsParameterPath(path string) bool {
	return isJSONInFolder(path, "parameters")
}

// IsProjectPath returns true if the given path describes a JSON file in the
// folder where data sets of type `Project` are stored.
func IsProjectPath(path string) bool {
	return isJSONInFolder(path, "projects")
}

// IsSocialIndicatorPath returns true if the given path describes a JSON file in the
// folder where data sets of type `SocialIndicator` are stored.
func IsSocialIndicatorPath(path string) bool {
	return isJSONInFolder(path, "social_indicators")
}

// Returns true if the given path describes a JSON file in the given folder.
func isJSONInFolder(path, folder string) bool {
	p := strings.ToLower(path)
	if !strings.Contains(p, folder) {
		return false
	}
	return strings.HasSuffix(p, ".json")
}
