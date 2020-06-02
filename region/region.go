package region

import (
	"github.com/thebrubaker/colony/colonist"
	"github.com/thebrubaker/colony/storable"
)

type Region struct {
	Colonists map[string]*colonist.Colonist `json:"colonists"`
	Stockpile *storable.Collection          `json:"stockpile"`
}
