package golcas

import (
	"encoding/json"
	"testing"
)

func TestActor(t *testing.T) {
	a := &Actor{
		ID:       "abc",
		Name:     "new actor",
		Category: "Category/bca/category"}
	bytes, err := json.Marshal(a)
	if err != nil {
		t.Error("failed to marshal Actor")
		return
	}
	clone := &Actor{}
	err = json.Unmarshal(bytes, clone)
	if err != nil || a.ID != clone.ID || a.Category != clone.Category {
		t.Error("failed to unmarshal Actor")
	}
	t.Log("Tested Actor:", string(bytes))
}

func TestSource(t *testing.T) {
	s := &Source{ID: "abc", Name: "Source"}
	bytes, err := json.Marshal(s)
	if err != nil {
		t.Error("failed to marshal Source")
		return
	}
	clone := &Source{}
	err = json.Unmarshal(bytes, clone)
	if err != nil || s.ID != clone.ID || s.Name != clone.Name {
		t.Error("failed to unmarshal Source")
		return
	}
	t.Log("Tested Source:", string(bytes))
}

func TestFlow(t *testing.T) {
	f := &Flow{ID: "abc", Name: "Flow", FlowType: PRODUCT_FLOW}
	bytes, _ := json.Marshal(f)
	clone := &Flow{}
	json.Unmarshal(bytes, clone)
	if clone.FlowType != PRODUCT_FLOW {
		t.Error("failed to unmarshal Flow")
		return
	}
	t.Log("Tested Flow:", string(bytes))
}
