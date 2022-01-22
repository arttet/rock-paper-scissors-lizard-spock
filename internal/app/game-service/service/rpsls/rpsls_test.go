package rpsls_test

import (
	"context"

	"github.com/arttet/rock-paper-scissors-lizard-spock/internal/app/game-service/service/rpsls"
	"github.com/arttet/rock-paper-scissors-lizard-spock/internal/mock"
	"github.com/arttet/rock-paper-scissors-lizard-spock/internal/model"

	"github.com/golang/mock/gomock"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Rock Paper Scissors Lizard Spock Game", func() {
	var (
		ctx context.Context

		ctrl                *gomock.Controller
		mockRandomGenerator *mock.MockRandomGenerator

		game rpsls.RockPaperScissorsLizardSpockGame
	)

	BeforeEach(func() {
		ctx = context.Background()

		ctrl = gomock.NewController(GinkgoT())
		Expect(ctrl).ShouldNot(BeNil())

		mockRandomGenerator = mock.NewMockRandomGenerator(ctrl)
		Expect(mockRandomGenerator).ShouldNot(BeNil())

		config := zap.NewDevelopmentConfig()
		config.Level.SetLevel(zapcore.PanicLevel)
		config.DisableCaller = true
		config.DisableStacktrace = true
		logger, _ := config.Build()
		defer logger.Sync()

		game = rpsls.NewRockPaperScissorsLizardSpockGame(mockRandomGenerator, logger)
	})

	// ////////////////////////////////////////////////////////////////////////

	It("gets all the choices correctly", Label("choices"), func() {
		Expect(game.GetChoices()).Should(Equal(rpsls.Choices))
	})

	It("gets a randomly generated choice", func() {
		mockRandomGenerator.EXPECT().GetRandomNumber(gomock.Any()).Return(-1).Times(1)
		Expect(game.GetChoice(ctx)).Should(BeElementOf(rpsls.Choices))
	})

	// ////////////////////////////////////////////////////////////////////////

	Describe("The game rules determines", func() {
		When("the player chooses something", func() {
			for winner, losers := range rpsls.Victories {
				winner, losers := winner, losers
				for _, loser := range losers {
					loser := loser

					It("should win", func() {
						mockRandomGenerator.EXPECT().GetRandomNumber(gomock.Any()).Return(int(loser - 1)).Times(1)
						Expect(game.PlayRound(ctx, winner)).Should(Equal(model.NewResult(winner, loser, rpsls.Win)))
					})

					It("should tie", func() {
						mockRandomGenerator.EXPECT().GetRandomNumber(gomock.Any()).Return(int(winner - 1)).Times(1)
						Expect(game.PlayRound(ctx, winner)).Should(Equal(model.NewResult(winner, winner, rpsls.Tie)))
					})

					It("should lose", func() {
						mockRandomGenerator.EXPECT().GetRandomNumber(gomock.Any()).Return(int(winner - 1)).Times(1)
						Expect(game.PlayRound(ctx, loser)).Should(Equal(model.NewResult(loser, winner, rpsls.Lose)))
					})
				}
			}
		})
	})
})
