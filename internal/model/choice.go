package model

import pb "github.com/arttet/rock-paper-scissors-lizard-spock/pkg/game-service/v1"

type Choice struct {
	ID   int32
	Name string
}

type Choices []*Choice

func ToMessage(choice *Choice) *pb.Choice {
	return &pb.Choice{
		Id:   choice.ID,
		Name: choice.Name,
	}
}

func FromMessage(choice *pb.Choice) *Choice {
	return &Choice{
		ID:   choice.Id,
		Name: choice.Name,
	}
}

func ToMessages(choices Choices) []*pb.Choice {
	result := make([]*pb.Choice, 0, len(choices))
	for _, choice := range choices {
		result = append(result, &pb.Choice{
			Id:   choice.ID,
			Name: choice.Name,
		})
	}

	return result
}

func FromMessages(choices []*pb.Choice) Choices {
	result := make(Choices, 0, len(choices))
	for _, choice := range choices {
		result = append(result, &Choice{
			ID:   choice.Id,
			Name: choice.Name,
		})
	}

	return result
}
