package resources

import (
	"testing"
)

type Stackable interface {
	IsStackable() bool
}

func TestResources(t *testing.T) {
	var items []interface{}

	items = append(items, Wood)

	if items[0] != Wood {
		t.Error("items does not contain wood")
	}

}
