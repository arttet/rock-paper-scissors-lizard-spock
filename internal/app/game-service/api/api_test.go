package api_test

import (
	"context"

	"github.com/arttet/rock-paper-scissors-lizard-spock/internal/app/game-service/api"
	"github.com/arttet/rock-paper-scissors-lizard-spock/internal/app/game-service/service"
	"github.com/arttet/rock-paper-scissors-lizard-spock/internal/model"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	pb "github.com/arttet/rock-paper-scissors-lizard-spock/pkg/game-service/v1"
)

var _ = Describe("Game Server", func() {
	var (
		err error
		ctx context.Context

		server pb.GameServiceServer
	)

	BeforeEach(func() {
		ctx = context.Background()

		config := zap.NewDevelopmentConfig()
		config.Level.SetLevel(zapcore.PanicLevel)
		config.DisableCaller = true
		config.DisableStacktrace = true
		logger, _ := config.Build()
		defer logger.Sync()

		server = api.NewGameServiceServer(logger)
	})

	// ////////////////////////////////////////////////////////////////////////

	Describe("gets all the choices that are usable for the UI", Label("GetChoicesV1", "choices"), func() {
		var (
			response *pb.GetChoicesV1Response
		)

		Context("when gets successfully", func() {
			BeforeEach(func() {
				response, err = server.GetChoicesV1(ctx, nil)
			})

			It("should not return an empty response", func() {
				Expect(response).ShouldNot(BeNil())
				Expect(err).Should(BeNil())

				Expect(model.FromMessages(response.Choice)).Should(HaveLen(5))
				Expect(model.FromMessages(response.Choice)).Should(BeEquivalentTo(service.GetChoices()))
			})
		})
	})

	// ////////////////////////////////////////////////////////////////////////

	Describe("gets a randomly generated choice", Label("GetChoiceV1", "choice"), func() {
		var (
			response *pb.GetChoiceV1Response
		)

		Context("when gets successfully", func() {
			BeforeEach(func() {
				response, err = server.GetChoiceV1(ctx, nil)
			})

			It("should not return an empty response", func() {
				Expect(response).ShouldNot(BeNil())
				Expect(err).Should(BeNil())

				Expect(model.FromMessage(response.Choice)).Should(BeElementOf(service.GetChoices()))
			})
		})
	})

	// ////////////////////////////////////////////////////////////////////////

	Describe("plays a round against a computer opponent", Label("PlayRoundV1", "play"), func() {
		var (
			request  *pb.PlayRoundV1Request
			response *pb.PlayRoundV1Response
		)

		Context("when gets successfully", func() {
			BeforeEach(func() {
				request = &pb.PlayRoundV1Request{
					Player: 1,
				}

				response, err = server.PlayRoundV1(ctx, request)
			})

			It("should not return an empty response", func() {
				Expect(response).ShouldNot(BeNil())
				Expect(err).Should(BeNil())
			})
		})
	})
})
