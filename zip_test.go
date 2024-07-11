package golcas

import (
	"os"
	"testing"
)

func TestPackIO(t *testing.T) {
	t.Log("create temporary file")
	file, _ := os.CreateTemp("", "test_schema_pack")
	path := file.Name()
	file.Close()
	if err := os.Remove(path); err != nil {
		t.Fatal("failed to delete temporary file")
	}

	t.Log("create zip file with actor abc", path)
	writer, _ := NewPackWriter(path)
	err := writer.WriteActor(&Actor{ID: "abc", Name: "My actor"})
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log("close zip file")
	if err := writer.Close(); err != nil {
		t.Fatal("failed to close zip writer")
	}

	t.Log("read actor from zip")
	reader, _ := NewZipReader(file.Name())
	actor, _ := reader.ReadActor("abc")
	if actor.Name != "My actor" {
		t.Fatal("Failed to get data set from package")
	}

	t.Log("iterate through actors in zip")
	found := false
	err = reader.EachActor(func(a *Actor) bool {
		if a.ID == "abc" {
			found = true
		}
		return true
	})
	if err != nil || !found {
		t.Fatal("failed to read actor")
	}
	t.Log("close zip reader")
	if err := reader.Close(); err != nil {
		t.Fatal("failed to close pack reader")
	}
	t.Log("delete zip file")
	if err := os.Remove(path); err != nil {
		t.Fatal("failed to delete zip")
	}
}
