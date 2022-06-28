package service

import (
	"Nexign/internal/model"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSpellerService_CheckOne(t *testing.T) {

	tests := []struct {
		name    string
		text    string
		want    []model.SpellerResponse
		wantErr bool
	}{
		{
			name:    "Нет ошибок",
			text:    "У зверей взрослеет смена",
			want:    []model.SpellerResponse{},
			wantErr: false,
		},
		{
			name:    "Ошибки",
			text:    "У зверей взраслеет смена",
			want:    []model.SpellerResponse{{Code: 1, Pos: 9, Row: 0, Word: "взраслеет", S: []string{"взрослеет"}}},
			wantErr: false,
		},
		{
			name:    "Невалид",
			text:    "У зверей 塞雷吉 смена",
			want:    []model.SpellerResponse{{Code: 4, Pos: 0, Row: 0, Word: "У зверей 塞雷吉 смена", S: []string{"У зверей 塞雷吉 смена"}}},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckOne(context.Background(), tt.text, &SpellerConfig{Url: "https://speller.yandex.net/services/spellservice.json/", Lang: "ru", Format: "plain"})
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestSpellerService_CheckMany(t *testing.T) {

	tests := []struct {
		name    string
		text    []string
		want    [][]model.SpellerResponse
		wantErr bool
	}{
		{
			name:    "Нет ошибок",
			text:    []string{"У зверей взрослеет смена", "Не обижайте лесных малышей"},
			want:    [][]model.SpellerResponse{[]model.SpellerResponse{}, []model.SpellerResponse{}},
			wantErr: false,
		},
		{
			name: "Ошибки",
			text: []string{"У зверей взраслеет смена", "Не обижайте лесных малашей"},
			want: [][]model.SpellerResponse{[]model.SpellerResponse{model.SpellerResponse{Code: 1, Pos: 9, Row: 0, Word: "взраслеет", S: []string{"взрослеет"}}},
				[]model.SpellerResponse{model.SpellerResponse{Code: 1, Pos: 19, Row: 0, Word: "малашей", S: []string{"малышей", "шалашей"}}}},
			wantErr: false,
		},
		{
			name:    "Невалид",
			text:    []string{"У зверей 塞雷吉 смена"},
			want:    [][]model.SpellerResponse{[]model.SpellerResponse{model.SpellerResponse{Code: 4, Pos: 0, Row: 0, Word: "У зверей 塞雷吉 смена", S: []string{"У зверей 塞雷吉 смена"}}}},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckMany(context.Background(), tt.text, &SpellerConfig{Url: "https://speller.yandex.net/services/spellservice.json/", Lang: "ru", Format: "plain"})
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
