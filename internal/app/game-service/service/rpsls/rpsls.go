package rpsls

import (
	"context"

	"github.com/arttet/rock-paper-scissors-lizard-spock/internal/app/game-service/service/random"
	"github.com/arttet/rock-paper-scissors-lizard-spock/internal/model"

	"go.uber.org/zap"
)

type RockPaperScissorsLizardSpockGame interface {
	GetChoices() model.Choices
	GetChoice(ctx context.Context) *model.Choice
	PlayRound(ctx context.Context, playerID int32) model.Result
}

func NewRockPaperScissorsLizardSpockGame(rg random.RandomGenerator, logger *zap.Logger) RockPaperScissorsLizardSpockGame {
	return &game{
		rg:     rg,
		logger: logger,
	}
}

type game struct {
	rg     random.RandomGenerator
	logger *zap.Logger
}

func (g *game) GetChoices() model.Choices {
	return Choices
}

func (g *game) GetChoice(ctx context.Context) *model.Choice {
	idx := g.rg.GetRandomNumber(ctx)
	if idx < 0 {
		idx = -idx
	}
	return Choices[idx%len(Choices)]
}

func (g *game) PlayRound(ctx context.Context, player int32) model.Result {
	computer := g.GetChoice(ctx)
	return model.NewResult(player, computer.ID, g.DetermineWinner(player, computer.ID))
}

func (g *game) DetermineWinner(playerID, computerID int32) string {
	if playerID == computerID {
		return Tie
	}

	defeats := Victories[playerID]
	if computerID == defeats[0] || computerID == defeats[1] {
		return Win
	}

	return Lose
}
