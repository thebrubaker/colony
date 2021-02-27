package region

import "github.com/thebrubaker/colony/storage"

type Region struct {
	Stockpile storage.Storage `json:"stockpile"`
}
