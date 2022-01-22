package internal

//go:generate mockgen --build_flags=--mod=mod -destination=./mock/rpsls_mock.go -package=mock github.com/arttet/rock-paper-scissors-lizard-spock/internal/app/game-service/service/rpsls RockPaperScissorsLizardSpockGame
//go:generate mockgen --build_flags=--mod=mod -destination=./mock/random_generator_mock.go -package=mock github.com/arttet/rock-paper-scissors-lizard-spock/internal/app/game-service/service/random RandomGenerator
