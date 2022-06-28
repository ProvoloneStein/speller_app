package handler

import (
	"Nexign/internal/model"
	"Nexign/internal/service"
	mock_service "Nexign/internal/service/mocks"
	"bytes"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestHandler_postText(t *testing.T) {
	// Init Test Table
	type mockBehavior func(s *mock_service.MockSpeller, ctx context.Context, text model.Speller)

	tests := []struct {
		name                 string
		inputBody            string
		inputStruct          model.Speller
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"text" : "У зверей взрослеет смена"}`,
			inputStruct: model.Speller{
				Text: "У зверей взрослеет смена",
			},
			mockBehavior: func(s *mock_service.MockSpeller, ctx context.Context, text model.Speller) {
				s.EXPECT().CreateOne(gomock.Any(), text).Return([]model.SpellerResponse{}, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"text":[]}`,
		},
		{
			name:                 "Wrong Input",
			inputBody:            `{"texts": "У зверей взраслеет смена"}`,
			inputStruct:          model.Speller{},
			mockBehavior:         func(s *mock_service.MockSpeller, ctx context.Context, text model.Speller) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid input body"}`,
		},
		{
			name:      "Service Error",
			inputBody: `{"text" : "У зверей взрослеет смена"}`,
			inputStruct: model.Speller{
				Text: "У зверей взрослеет смена",
			},
			mockBehavior: func(s *mock_service.MockSpeller, ctx context.Context, text model.Speller) {
				s.EXPECT().CreateOne(gomock.Any(), text).Return([]model.SpellerResponse{}, errors.New("Что-то пошло не так..."))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"Что-то пошло не так..."}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			cont := mock_service.NewMockSpeller(c)
			services := &service.Service{Speller: cont}
			handler := Handler{services}

			// Init Endpoint
			r := gin.New()
			r.POST("/api/v1/text", handler.postText)
			// Create Request
			w := httptest.NewRecorder()
			contxt, _ := gin.CreateTestContext(httptest.NewRecorder())
			test.mockBehavior(cont, contxt, test.inputStruct)
			req := httptest.NewRequest("POST", "/api/v1/text",
				bytes.NewBufferString(test.inputBody))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestHandler_postMany(t *testing.T) {
	// Init Test Table
	type mockBehavior func(s *mock_service.MockSpeller, ctx context.Context, text model.Spellers)

	tests := []struct {
		name                 string
		inputBody            string
		inputStruct          model.Spellers
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"texts" : ["У зверей взрослеет смена", "У зверей взрослеет смена"]}`,
			inputStruct: model.Spellers{
				Text: []string{"У зверей взрослеет смена", "У зверей взрослеет смена"},
			},
			mockBehavior: func(s *mock_service.MockSpeller, ctx context.Context, text model.Spellers) {
				s.EXPECT().CreateMany(gomock.Any(), text).Return([][]model.SpellerResponse{}, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"texts":[]}`,
		},
		{
			name:                 "Wrong Input",
			inputBody:            `{"text": "У зверей взраслеет смена"}`,
			inputStruct:          model.Spellers{},
			mockBehavior:         func(s *mock_service.MockSpeller, ctx context.Context, text model.Spellers) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid input body"}`,
		},
		{
			name:      "Service Error",
			inputBody: `{"texts" : ["У зверей взрослеет смена", "У зверей взрослеет смена"]}`,
			inputStruct: model.Spellers{
				Text: []string{"У зверей взрослеет смена", "У зверей взрослеет смена"},
			},
			mockBehavior: func(s *mock_service.MockSpeller, ctx context.Context, text model.Spellers) {
				s.EXPECT().CreateMany(gomock.Any(), text).Return([][]model.SpellerResponse{}, errors.New("Что-то пошло не так..."))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"Что-то пошло не так..."}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			cont := mock_service.NewMockSpeller(c)
			services := &service.Service{Speller: cont}
			handler := Handler{services}

			// Init Endpoint
			r := gin.New()
			r.POST("/api/v1/texts", handler.postMany)
			// Create Request
			w := httptest.NewRecorder()
			contxt, _ := gin.CreateTestContext(httptest.NewRecorder())
			test.mockBehavior(cont, contxt, test.inputStruct)
			req := httptest.NewRequest("POST", "/api/v1/texts",
				bytes.NewBufferString(test.inputBody))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
