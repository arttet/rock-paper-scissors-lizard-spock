package api_test

import (
	"context"
	"encoding/json"

	"github.com/arttet/rock-paper-scissors-lizard-spock/internal/app/game-service/api"
	"github.com/arttet/rock-paper-scissors-lizard-spock/internal/app/game-service/service/rpsls"
	"github.com/arttet/rock-paper-scissors-lizard-spock/internal/mock"
	"github.com/arttet/rock-paper-scissors-lizard-spock/internal/model"

	"google.golang.org/genproto/googleapis/api/httpbody"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/golang/mock/gomock"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	pb "github.com/arttet/rock-paper-scissors-lizard-spock/pkg/game-service/v1"
)

var _ = Describe("API Server", func() {
	var (
		err error
		ctx context.Context

		ctrl     *gomock.Controller
		mockGame *mock.MockRockPaperScissorsLizardSpockGame

		server pb.GameServiceServer
	)

	BeforeEach(func() {
		ctx = context.Background()

		ctrl = gomock.NewController(GinkgoT())
		Expect(ctrl).ShouldNot(BeNil())

		mockGame = mock.NewMockRockPaperScissorsLizardSpockGame(ctrl)
		Expect(mockGame).ShouldNot(BeNil())

		config := zap.NewDevelopmentConfig()
		config.Level.SetLevel(zapcore.PanicLevel)
		config.DisableCaller = true
		config.DisableStacktrace = true
		logger, _ := config.Build()
		defer logger.Sync()

		server = api.NewGameServiceServer(mockGame, logger)
	})

	// ////////////////////////////////////////////////////////////////////////

	Describe("gets all the choices", Label("GetChoicesV1", "choices"), func() {
		var (
			response *httpbody.HttpBody
		)

		Context("when a successful request", func() {
			BeforeEach(func() {
				mockGame.EXPECT().GetChoices().Return(rpsls.Choices).Times(1)
				response, err = server.GetChoicesV1(ctx, nil)
			})

			It("should return correct data", func() {
				Expect(response).ShouldNot(BeNil())
				Expect(err).Should(BeNil())

				Expect(response.ContentType).Should(Equal("application/json"))

				var choices model.Choices
				err = json.Unmarshal(response.Data, &choices)
				Expect(choices).Should(HaveLen(5))
				Expect(choices).Should(BeEquivalentTo(rpsls.Choices))
			})
		})
	})

	// ////////////////////////////////////////////////////////////////////////

	Describe("gets a randomly generated choice", Label("GetChoiceV1", "choice"), func() {
		for _, choice := range rpsls.Choices {
			var (
				choice = choice

				response *pb.GetChoiceV1Response
			)

			Context("when a successful request", func() {
				BeforeEach(func() {
					mockGame.EXPECT().GetChoice(gomock.Any()).Return(choice).Times(1)
					response, err = server.GetChoiceV1(ctx, nil)
				})

				It("should return a randomly generated choice", func() {
					Expect(response).ShouldNot(BeNil())
					Expect(err).Should(BeNil())

					Expect(response.Id).Should(Equal(choice.ID))
					Expect(response.Name).Should(Equal(choice.Name))
				})
			})
		}
	})

	// ////////////////////////////////////////////////////////////////////////

	Describe("gets results", Label("PlayRoundV1", "play"), func() {
		for _, choice := range rpsls.Choices {
			var (
				choice = choice

				request  *pb.PlayRoundV1Request
				response *pb.PlayRoundV1Response

				mockResult = model.Result{
					Player:   choice.ID,
					Computer: choice.ID,
					Results:  rpsls.Tie,
				}
			)

			Context("when a successful request", func() {
				BeforeEach(func() {
					mockGame.EXPECT().PlayRound(gomock.Any(), gomock.Any()).Return(mockResult).Times(1)
					request = &pb.PlayRoundV1Request{Player: choice.ID}
					response, err = server.PlayRoundV1(ctx, request)
				})

				It("should determine a winner", func() {
					Expect(response).ShouldNot(BeNil())
					Expect(err).Should(BeNil())

					Expect(response.Player).Should(Equal(choice.ID))
					Expect(response.Computer).Should(Equal(choice.ID))
					Expect(response.Results).Should(Equal(rpsls.Tie))
				})
			})
		}
	})
})
