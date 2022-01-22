package model

type Choice struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type Choices []*Choice
