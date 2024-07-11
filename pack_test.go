package golcas

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestPackIO(t *testing.T) {
	//actor := NewActor("abc", "My actor")
	file, _ := ioutil.TempFile("", "test_schema_pack")
	path := file.Name()
	os.Remove(path)
	writer, _ := NewPackWriter(path)
	err := writer.PutActor(&Actor{ID: "abc", Name: "My actor"})
	if err != nil {
		t.Error(err)
		return
	}
	writer.Close()
	t.Log("created zip file", path)
	reader, _ := NewPackReader(file.Name())
	actor, _ := reader.GetActor("abc")
	if actor.Name != "My actor" {
		t.Error("Failed to get data set from package")
	}
	reader.Close()
	os.Remove(path)
}
