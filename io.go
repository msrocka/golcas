package ld

import (
	"encoding/json"
)

// ReadCategory converts the given bytes into a Category
func ReadCategory(data []byte) (*Category, error) {
	v := &Category{}
	if err := json.Unmarshal(data, v); err != nil {
		return nil, err
	}
	return v, nil
}

// ReadActor converts the given bytes into a Actor
func ReadActor(data []byte) (*Actor, error) {
	v := &Actor{}
	if err := json.Unmarshal(data, v); err != nil {
		return nil, err
	}
	return v, nil
}

// ReadSource converts the given bytes into a Source
func ReadSource(data []byte) (*Source, error) {
	v := &Source{}
	if err := json.Unmarshal(data, v); err != nil {
		return nil, err
	}
	return v, nil
}

// ReadUnitGroup converts the given bytes into a UnitGroup
func ReadUnitGroup(data []byte) (*UnitGroup, error) {
	v := &UnitGroup{}
	if err := json.Unmarshal(data, v); err != nil {
		return nil, err
	}
	return v, nil
}

// ReadFlowProperty converts the given bytes into a FlowProperty
func ReadFlowProperty(data []byte) (*FlowProperty, error) {
	v := &FlowProperty{}
	if err := json.Unmarshal(data, v); err != nil {
		return nil, err
	}
	return v, nil
}

// ReadFlow converts the given bytes into a Flow
func ReadFlow(data []byte) (*Flow, error) {
	v := &Flow{}
	if err := json.Unmarshal(data, v); err != nil {
		return nil, err
	}
	return v, nil
}

// ReadProcess converts the given bytes into a Process
func ReadProcess(data []byte) (*Process, error) {
	v := &Process{}
	if err := json.Unmarshal(data, v); err != nil {
		return nil, err
	}
	return v, nil
}

// ReadImpactCategory converts the given bytes into a ImpactCategory
func ReadImpactCategory(data []byte) (*ImpactCategory, error) {
	v := &ImpactCategory{}
	if err := json.Unmarshal(data, v); err != nil {
		return nil, err
	}
	return v, nil
}

// ReadImpactMethod converts the given bytes into a ImpactMethod
func ReadImpactMethod(data []byte) (*ImpactMethod, error) {
	v := &ImpactMethod{}
	if err := json.Unmarshal(data, v); err != nil {
		return nil, err
	}
	return v, nil
}
