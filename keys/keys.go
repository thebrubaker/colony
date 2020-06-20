package keys

import "github.com/rs/xid"

type GameKey string
type ColonistKey string

func NewGameKey() GameKey {
	return GameKey(xid.New().String())
}
