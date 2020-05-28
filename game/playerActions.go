package game

type PlayerActions struct {
	Queue []*PlayerAction
}

type PlayerAction struct {
	TypeID string
}
