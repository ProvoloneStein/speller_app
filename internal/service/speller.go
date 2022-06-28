package service

import (
	"Nexign/internal/model"
	"Nexign/pkg/logger"
	"context"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"net/url"
)

type SpellerService struct {
	cfg *SpellerConfig
}

func NewSpellerService(cfg *SpellerConfig) *SpellerService {
	return &SpellerService{cfg: cfg}
}

// Check
func CheckOne(ctx context.Context, text string, cfg *SpellerConfig) ([]model.SpellerResponse, error) {
	var words []model.SpellerResponse
	serviceMethod := "checkText"
	resp, err := http.PostForm(cfg.Url + serviceMethod, url.Values{
		"text":   {text},
		"lang":   {cfg.Lang},
		"format": {cfg.Format},
	})
	if err != nil {
		logger.GetLogger(ctx).Error("Spellservice error:", zap.Error(err))
		return words, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.GetLogger(ctx).Error("Spellservice error:", zap.Error(err))
		return words, err
	}

	err = resp.Body.Close()
	if err != nil {
		logger.GetLogger(ctx).Error("Spellservice error:", zap.Error(err))
	}

	if err = json.Unmarshal(body, &words); err != nil {
		logger.GetLogger(ctx).Error("Spellservice error:", zap.Error(err))
		return words, err
	}
	fmt.Println(words)

	return words, nil
}

func (s *SpellerService) CreateOne(ctx context.Context, text model.Speller) ([]model.SpellerResponse, error) {
	return CheckOne(ctx, text.Text, s.cfg)
}

//
func CheckMany(ctx context.Context, text []string, cfg *SpellerConfig) ([][]model.SpellerResponse, error) {
	var words [][]model.SpellerResponse
	serviceMethod := "checkTexts"
	resp, err := http.PostForm(cfg.Url + serviceMethod, url.Values{
		"text":   text,
		"lang":   {cfg.Lang},
		"format": {cfg.Format},
	})
	if err != nil {
		logger.GetLogger(ctx).Error("Spellservice error:", zap.Error(err))
		return words, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.GetLogger(ctx).Error("Spellservice error:", zap.Error(err))
		return words, err
	}

	err = resp.Body.Close()
	if err != nil {
		logger.GetLogger(ctx).Error("Spellservice error:", zap.Error(err))
	}
	if err = json.Unmarshal(body, &words); err != nil {
		logger.GetLogger(ctx).Error("Spellservice error:", zap.Error(err))
		return words, err
	}
	return words, nil
}

func (s *SpellerService) CreateMany(ctx context.Context, texts model.Spellers) ([][]model.SpellerResponse, error) {
	return CheckMany(ctx, texts.Text, s.cfg)
}
