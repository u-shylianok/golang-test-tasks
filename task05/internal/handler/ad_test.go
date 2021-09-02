package handler

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/u-shylianok/golang-test-tasks/task05/internal/model"
	"github.com/u-shylianok/golang-test-tasks/task05/internal/service"
	mock_service "github.com/u-shylianok/golang-test-tasks/task05/internal/service/mocks"
)

func TestHandler_createAd(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAd, ad model.Ad)

	times := make(map[string]time.Time)
	times["OK"], _ = time.Parse("2006-01-02T15:04:05Z", "2021-08-30T19:17:25Z")
	description := "description"

	testTable := []struct {
		name                string
		inputBody           string
		inputAd             model.Ad
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name: "OK",
			inputBody: `{
				"name": "name",
				"date": "2021-08-30T19:17:25Z",
				"price": 10,
				"description": "description",
				"photos": [
					{
						"link": "link"
					},
					{
						"link": "link"
					},
					{
						"link": "link"
					}
				]
			}`,
			inputAd: model.Ad{
				Name:        "name",
				Date:        times["OK"],
				Price:       10,
				Description: &description,
				Photos: &[]model.Photo{
					{Link: "link"},
					{Link: "link"},
					{Link: "link"},
				},
			},
			mockBehavior: func(s *mock_service.MockAd, ad model.Ad) {
				s.EXPECT().Create(ad).Return(1, nil)
			},
			expectedStatusCode:  http.StatusOK,
			expectedRequestBody: `{"id":1}`,
		},
		{
			name: "Empty Fields 1",
			inputBody: `{
				"name": "name",
				"date": "2021-08-30T19:17:25Z"
			}`,
			inputAd:             model.Ad{},
			mockBehavior:        func(s *mock_service.MockAd, ad model.Ad) {},
			expectedStatusCode:  http.StatusBadRequest,
			expectedRequestBody: `{"message":"invalid input body"}`,
		},
		{
			name: "Empty Fields 2",
			inputBody: `{
				"name": "name",
				"date": "2021-08-30T19:17:25Z",
				"description": "description",
				"photos": [
					{
						"link": "link"
					},
					{
						"link": "link"
					},
					{
						"link": "link"
					}
				]
			}`,
			inputAd:             model.Ad{},
			mockBehavior:        func(s *mock_service.MockAd, ad model.Ad) {},
			expectedStatusCode:  http.StatusBadRequest,
			expectedRequestBody: `{"message":"invalid input body"}`,
		},
		{
			name: "Name 200",
			inputBody: `{
				"name": "namenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamenamename",
				"date": "2021-08-30T19:17:25Z",
				"price": 10,
				"description": "description",
				"photos": [
					{
						"link": "link"
					},
					{
						"link": "link"
					},
					{
						"link": "link"
					}
				]
			}`,
			inputAd:             model.Ad{},
			mockBehavior:        func(s *mock_service.MockAd, ad model.Ad) {},
			expectedStatusCode:  http.StatusBadRequest,
			expectedRequestBody: `{"message":"name should be no more than 200 symbols"}`,
		},
		{
			name: "Description 1000",
			inputBody: `{
				"name": "name",
				"date": "2021-08-30T19:17:25Z",
				"price": 10,
				"description": "descriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescription",
				"photos": [
					{
						"link": "link"
					},
					{
						"link": "link"
					},
					{
						"link": "link"
					}
				]
			}`,
			inputAd:             model.Ad{},
			mockBehavior:        func(s *mock_service.MockAd, ad model.Ad) {},
			expectedStatusCode:  http.StatusBadRequest,
			expectedRequestBody: `{"message":"description should be no more than 1000 symbols"}`,
		},
		{
			name: "Photos 1",
			inputBody: `{
				"name": "name",
				"date": "2021-08-30T19:17:25Z",
				"price": 10,
				"description": "description"
			}`,
			inputAd:             model.Ad{},
			mockBehavior:        func(s *mock_service.MockAd, ad model.Ad) {},
			expectedStatusCode:  http.StatusBadRequest,
			expectedRequestBody: `{"message":"must be at least 1 photo"}`,
		},
		{
			name: "Photos 2",
			inputBody: `{
				"name": "name",
				"date": "2021-08-30T19:17:25Z",
				"price": 10,
				"description": "description",
				"photos": [
					{
						"link": "link"
					},
					{
						"link": "link"
					},
					{
						"link": "link"
					},
					{
						"link": "link"
					}
				]
			}`,
			inputAd:             model.Ad{},
			mockBehavior:        func(s *mock_service.MockAd, ad model.Ad) {},
			expectedStatusCode:  http.StatusBadRequest,
			expectedRequestBody: `{"message":"should be no more than 3 photos"}`,
		},
		{
			name: "Service Failure",
			inputBody: `{
				"name": "name",
				"date": "2021-08-30T19:17:25Z",
				"price": 10,
				"description": "description",
				"photos": [
					{
						"link": "link"
					},
					{
						"link": "link"
					},
					{
						"link": "link"
					}
				]
			}`,
			inputAd: model.Ad{
				Name:        "name",
				Date:        times["OK"],
				Price:       10,
				Description: &description,
				Photos: &[]model.Photo{
					{Link: "link"},
					{Link: "link"},
					{Link: "link"},
				},
			},
			mockBehavior: func(s *mock_service.MockAd, ad model.Ad) {
				s.EXPECT().Create(ad).Return(1, errors.New("service failure"))
			},
			expectedStatusCode:  http.StatusInternalServerError,
			expectedRequestBody: `{"message":"service failure"}`,
		},
	}

	gin.SetMode(gin.ReleaseMode)
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// init deps
			c := gomock.NewController(t)
			defer c.Finish()

			ad := mock_service.NewMockAd(c)
			testCase.mockBehavior(ad, testCase.inputAd)

			services := &service.Service{Ad: ad}
			handler := NewHandler(services)

			// test server
			r := gin.New()
			r.POST("/", handler.createAd)

			// test request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/", bytes.NewBufferString(testCase.inputBody))

			// perform request
			r.ServeHTTP(w, req)

			// assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())

		})
	}
}
