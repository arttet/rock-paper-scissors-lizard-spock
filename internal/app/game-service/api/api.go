package api

import (
	"context"
	"encoding/json"

	"github.com/arttet/rock-paper-scissors-lizard-spock/internal/app/game-service/service/rpsls"

	"go.uber.org/zap"
	"google.golang.org/genproto/googleapis/api/httpbody"

	"github.com/golang/protobuf/ptypes/empty"

	pb "github.com/arttet/rock-paper-scissors-lizard-spock/pkg/game-service/v1"
)

type api struct {
	pb.UnimplementedGameServiceServer
	game   rpsls.RockPaperScissorsLizardSpockGame
	logger *zap.Logger
}

func NewGameServiceServer(
	game rpsls.RockPaperScissorsLizardSpockGame,
	logger *zap.Logger,
) pb.GameServiceServer {

	return &api{
		game:   game,
		logger: logger,
	}
}

func (a *api) GetChoicesV1(
	ctx context.Context,
	_ *empty.Empty,
) (
	*httpbody.HttpBody,
	error,
) {

	data, _ := json.Marshal(a.game.GetChoices()) // nolint:errcheck

	return &httpbody.HttpBody{
		ContentType: "application/json",
		Data:        data,
	}, nil
}

func (a *api) GetChoiceV1(
	ctx context.Context,
	_ *empty.Empty,
) (
	*pb.GetChoiceV1Response,
	error,
) {

	choice := a.game.GetChoice(ctx)

	return &pb.GetChoiceV1Response{
		Id:   choice.ID,
		Name: choice.Name,
	}, nil
}

func (a *api) PlayRoundV1(
	ctx context.Context,
	request *pb.PlayRoundV1Request,
) (
	*pb.PlayRoundV1Response,
	error,
) {

	result := a.game.PlayRound(ctx, request.Player)

	return &pb.PlayRoundV1Response{
		Player:   result.Player,
		Computer: result.Computer,
		Results:  result.Results,
	}, nil
}
