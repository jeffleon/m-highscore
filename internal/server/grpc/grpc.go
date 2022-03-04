package grpc

import (
	"context"

	pbhighscore "github.com/jeffleon/m-apis/tree/master/m-highscore/v1"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

type Grpc struct {
	address string
	srv     *grpc.Server
}

var HighScore = 99999999999.0

func (g *Grpc) SetHighScore(ctx context.Context, input *pbhighscore.SetHighScoreRequest) (*pbhighscore.SethighScoreResponse, error) {
	log.Info().Msg("Setting in highscore is called")
	HighScore = input.HighScore
	return &pbhighscore.SethighScoreResponse{
		Set: true,
	}, nil
}

func (g *Grpc) GetHighScore(ctx context.Context, input *pbhighscore.GetHighScoreRequest) (*pbhighscore.GethighScoreResponse, error) {
	log.Info().Msg("Getting in highscore is called")
	return &pbhighscore.SethighScoreResponse{
		HighScore: HighScore,
	}, nil
}
