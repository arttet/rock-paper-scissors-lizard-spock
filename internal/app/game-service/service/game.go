package service

import (
	"math/rand"
	"time"

	"github.com/arttet/rock-paper-scissors-lizard-spock/internal/model"
)

var choices model.Choices = model.Choices{
	&model.Choice{
		ID:   1,
		Name: "rock",
	},
	&model.Choice{
		ID:   2,
		Name: "paper",
	},
	&model.Choice{
		ID:   3,
		Name: "scissors",
	},
	&model.Choice{
		ID:   4,
		Name: "lizard",
	},
	&model.Choice{
		ID:   5,
		Name: "spock",
	},
}

func GetChoices() model.Choices {
	return choices
}

func GetRandomChoice(seed int64) *model.Choice {
	if seed == 0 {
		seed = time.Now().UnixNano()
	}

	rand.Seed(seed)
	idx := rand.Intn(len(choices)) // nolint:gosec

	return choices[idx]
}
