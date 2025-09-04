package storage

import (
	"testing"
)

func TestOpen(t *testing.T) {
	file, err := DefaultFileStorage.open()
	if err != nil {
		t.Error(err)
	}
	name := file.Name()
	t.Logf("file name: %s", name)
}
