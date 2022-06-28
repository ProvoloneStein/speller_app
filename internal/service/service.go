package service

import (
	"Nexign/internal/model"
	"context"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type SpellerConfig struct {
	Url string
	Lang string
	Format string
}

type Speller interface {
	CreateOne(ctx context.Context, text model.Speller) ([]model.SpellerResponse, error)
	CreateMany(ctx context.Context, texts model.Spellers) ([][]model.SpellerResponse, error)
}

type Service struct {
	Speller
}

func NewService(cfg *SpellerConfig) *Service {
	return &Service{
		Speller: NewSpellerService(cfg),
	}
}
