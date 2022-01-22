package random

import (
	"context"
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"time"

	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

type RandomGenerator interface {
	GetRandomNumber(ctx context.Context) int
}

func NewRandomGenerator(n int, seed int64, logger *zap.Logger) RandomGenerator {
	return &rg{
		n:      n,
		seed:   seed,
		logger: logger,
	}
}

type rg struct {
	RandomNumber int `json:"random_number"` // nolint:tagliatelle
	seed         int64
	n            int
	logger       *zap.Logger
}

func (r *rg) GetRandomNumber(ctx context.Context) int {
	span, _ := opentracing.StartSpanFromContext(ctx, "GetRandomNumber")
	defer span.Finish()

	body, err := MakeRequest(ctx, http.MethodGet, "https://codechallenge.boohma.com/random", nil)
	if err != nil {
		r.logger.Error("failed to get a random number", zap.Error(err))
		return r.GenerateRandomNumber()
	}

	err = json.Unmarshal(body, r)
	if err != nil {
		r.logger.Error("failed to unmarshal data", zap.Error(err))
		return r.GenerateRandomNumber()
	}

	return r.RandomNumber
}

func (r *rg) GenerateRandomNumber() int {
	seed := r.seed
	if seed == 0 {
		seed = time.Now().UnixNano()
	}
	rand.Seed(seed)
	return rand.Intn(r.n) // nolint:gosec
}

func MakeRequest(
	ctx context.Context,
	method string,
	url string,
	body io.Reader,
) (
	[]byte,
	error,
) {

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
