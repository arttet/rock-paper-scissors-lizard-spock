package model

type Result struct {
	Player   int32
	Computer int32
	Results  string
}

func NewResult(player, computer int32, results string) Result {
	return Result{
		Player:   player,
		Computer: computer,
		Results:  results,
	}
}
