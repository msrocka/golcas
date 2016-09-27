package schema

import (
	"encoding/json"
	"testing"
)

func TestCategory(t *testing.T) {
	c := NewCategory("abc", "category", ProcessModel)
	c.Category = NewReference("Category", "bca", "parent category")
	bytes, err := json.Marshal(c)
	if err != nil {
		t.Error("failed to marshal Category")
		return
	}
	clone := &Category{}
	err = json.Unmarshal(bytes, clone)
	if err != nil || c.ID != clone.ID || c.Name != clone.Name {
		t.Error("failed to unmarshal Category")
		return
	}
	if c.Category.ID != clone.Category.ID {
		t.Error("failed to unmarshal parent reference")
	}
	t.Log("Tested category: \n", string(bytes))
}

func TestActor(t *testing.T) {
	a := NewActor("abc", "new actor")
	a.Category = NewReference("Category", "bca", "category")
	bytes, err := json.Marshal(a)
	if err != nil {
		t.Error("failed to marshal Actor")
		return
	}
	clone := &Actor{}
	err = json.Unmarshal(bytes, clone)
	if err != nil || a.ID != clone.ID || a.Category.ID != clone.Category.ID {
		t.Error("failed to unmarshal Actor")
	}
	t.Log("Tested Actor:", string(bytes))
}

func TestSource(t *testing.T) {
	s := NewSource("abc", "Source")
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
	f := NewFlow("abc", "Flow")
	f.Type = ProductFlow
	bytes, _ := json.Marshal(f)
	clone := &Flow{}
	json.Unmarshal(bytes, clone)
	if clone.Type != ProductFlow {
		t.Error("failed to unmarshal Flow")
		return
	}
	t.Log("Tested Flow:", string(bytes))
}
