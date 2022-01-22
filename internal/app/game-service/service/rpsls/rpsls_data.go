package rpsls

import (
	"github.com/arttet/rock-paper-scissors-lizard-spock/internal/model"
)

const (
	Rock int32 = iota + 1
	Paper
	Scissors
	Lizard
	Spock
)

const (
	Win  = "win"
	Lose = "lose"
	Tie  = "tie"
)

var Choices = model.Choices{
	&model.Choice{
		ID:   Rock,
		Name: "rock",
	},
	&model.Choice{
		ID:   Paper,
		Name: "paper",
	},
	&model.Choice{
		ID:   Scissors,
		Name: "scissors",
	},
	&model.Choice{
		ID:   Lizard,
		Name: "lizard",
	},
	&model.Choice{
		ID:   Spock,
		Name: "spock",
	},
}

var Victories = map[int32][2]int32{
	Rock:     {Scissors, Lizard},
	Paper:    {Rock, Spock},
	Scissors: {Paper, Lizard},
	Lizard:   {Paper, Spock},
	Spock:    {Scissors, Rock},
}
