package golcas

import (
	"encoding/json"
)

func ReadActor(data []byte) (*Actor, error) {
	v := &Actor{}
	if err := json.Unmarshal(data, v); err != nil {
		return nil, err
	}
	return v, nil
}

func ReadCurrency(data []byte) (*Currency, error) {
	v := &Currency{}
	if err := json.Unmarshal(data, v); err != nil {
		return nil, err
	}
	return v, nil
}

func ReadDQSystem(data []byte) (*DQSystem, error) {
	v := &DQSystem{}
	if err := json.Unmarshal(data, v); err != nil {
		return nil, err
	}
	return v, nil
}

func ReadEpd(data []byte) (*Epd, error) {
	v := &Epd{}
	if err := json.Unmarshal(data, v); err != nil {
		return nil, err
	}
	return v, nil
}

func ReadFlow(data []byte) (*Flow, error) {
	v := &Flow{}
	if err := json.Unmarshal(data, v); err != nil {
		return nil, err
	}
	return v, nil
}

func ReadFlowProperty(data []byte) (*FlowProperty, error) {
	v := &FlowProperty{}
	if err := json.Unmarshal(data, v); err != nil {
		return nil, err
	}
	return v, nil
}

func ReadImpactCategory(data []byte) (*ImpactCategory, error) {
	v := &ImpactCategory{}
	if err := json.Unmarshal(data, v); err != nil {
		return nil, err
	}
	return v, nil
}

func ReadImpactMethod(data []byte) (*ImpactMethod, error) {
	v := &ImpactMethod{}
	if err := json.Unmarshal(data, v); err != nil {
		return nil, err
	}
	return v, nil
}

func ReadLocation(data []byte) (*Location, error) {
	v := &Location{}
	if err := json.Unmarshal(data, v); err != nil {
		return nil, err
	}
	return v, nil
}

func ReadParameter(data []byte) (*Parameter, error) {
	v := &Parameter{}
	if err := json.Unmarshal(data, v); err != nil {
		return nil, err
	}
	return v, nil
}

func ReadProcess(data []byte) (*Process, error) {
	v := &Process{}
	if err := json.Unmarshal(data, v); err != nil {
		return nil, err
	}
	return v, nil
}

func ReadProductSystem(data []byte) (*ProductSystem, error) {
	v := &ProductSystem{}
	if err := json.Unmarshal(data, v); err != nil {
		return nil, err
	}
	return v, nil
}

func ReadProject(data []byte) (*Project, error) {
	v := &Project{}
	if err := json.Unmarshal(data, v); err != nil {
		return nil, err
	}
	return v, nil
}

func ReadResult(data []byte) (*Result, error) {
	v := &Result{}
	if err := json.Unmarshal(data, v); err != nil {
		return nil, err
	}
	return v, nil
}

func ReadSocialIndicator(data []byte) (*SocialIndicator, error) {
	v := &SocialIndicator{}
	if err := json.Unmarshal(data, v); err != nil {
		return nil, err
	}
	return v, nil
}

func ReadSource(data []byte) (*Source, error) {
	v := &Source{}
	if err := json.Unmarshal(data, v); err != nil {
		return nil, err
	}
	return v, nil
}

func ReadUnitGroup(data []byte) (*UnitGroup, error) {
	v := &UnitGroup{}
	if err := json.Unmarshal(data, v); err != nil {
		return nil, err
	}
	return v, nil
}
