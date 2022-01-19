package api

import (
	"context"

	"github.com/arttet/rock-paper-scissors-lizard-spock/internal/app/game-service/service"
	"github.com/arttet/rock-paper-scissors-lizard-spock/internal/model"

	"go.uber.org/zap"

	"github.com/golang/protobuf/ptypes/empty"

	pb "github.com/arttet/rock-paper-scissors-lizard-spock/pkg/game-service/v1"
)

type api struct {
	pb.UnimplementedGameServiceServer
	logger *zap.Logger
}

func NewGameServiceServer(logger *zap.Logger) pb.GameServiceServer {
	return &api{
		logger: logger,
	}
}

func (a *api) GetChoicesV1(
	ctx context.Context,
	_ *empty.Empty,
) (
	*pb.GetChoicesV1Response,
	error,
) {

	choices := service.GetChoices()

	return &pb.GetChoicesV1Response{
		Choice: model.ToMessages(choices),
	}, nil
}

func (a *api) GetChoiceV1(
	ctx context.Context,
	_ *empty.Empty,
) (
	*pb.GetChoiceV1Response,
	error,
) {

	choice := service.GetRandomChoice(0)

	return &pb.GetChoiceV1Response{
		Choice: model.ToMessage(choice),
	}, nil
}

func (a *api) PlayRoundV1(
	ctx context.Context,
	request *pb.PlayRoundV1Request,
) (
	*pb.PlayRoundV1Response,
	error,
) {

	return &pb.PlayRoundV1Response{
		Player:   1,
		Computer: 1,
		Results:  "win",
	}, nil
}
