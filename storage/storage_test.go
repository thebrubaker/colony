package storage

import (
	"testing"

	"github.com/thebrubaker/colony/resources"
)

func TestStorage(t *testing.T) {
	storage := Storage{
		Size: 10,
	}
	var err error

	err = storage.Add(resources.Wood, 10)

	if err != nil {
		t.Error(err)
	}

	if !storage.Has(resources.Wood, 10) {
		t.Error("wood was not placed in storage")
	}

	err = storage.Remove(resources.Wood, 5)

	if err != nil {
		t.Error(err)
	}

	if !storage.Has(resources.Wood, 5) {
		t.Error("wood was not removed from storage")
	}
}
