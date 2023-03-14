package handler

import (
	"bytes"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	mementor_back "mementor-back"
	"mementor-back/pkg/service"
	mock_service "mementor-back/pkg/service/mocks"
	"net/http/httptest"
	"testing"
)

func TestHandlerSignUp(t *testing.T) {
	data := []struct {
		test_name             string
		input_body            string
		input_user            mementor_back.Auth
		mock_func             func(s *mock_service.MockAuthorization, auth mementor_back.Auth)
		content_type          string
		expected_status_code  int
		expected_request_body string
	}{{
		test_name:  "good",
		input_body: `{"email":"123@withoutasecondthought.com", "password":"passed"}`,
		input_user: mementor_back.Auth{
			Email:    "123@withoutasecondthought.com",
			Password: "passed",
		},
		mock_func: func(s *mock_service.MockAuthorization, auth mementor_back.Auth) {
			ctx := context.Background()
			s.EXPECT().SignUp(ctx, auth).Return("1", nil)
		},
		content_type:          "application/json",
		expected_status_code:  200,
		expected_request_body: `{"token":"1"}`,
	},
		{
			test_name:  "empty fields",
			input_body: `{"email": "123@withoutasecondthought.com", "password":""}`,
			input_user: mementor_back.Auth{
				Email:    "test",
				Password: "",
			},
			mock_func:             func(s *mock_service.MockAuthorization, auth mementor_back.Auth) {},
			content_type:          "application/json",
			expected_status_code:  400,
			expected_request_body: `{"message":"validation error"}`,
		},
		{
			test_name:  "email error",
			input_body: `{"email": "@withoutasecondthought", "password":"123321"}`,
			input_user: mementor_back.Auth{
				Email:    "test",
				Password: "",
			},
			mock_func:             func(s *mock_service.MockAuthorization, auth mementor_back.Auth) {},
			content_type:          "application/json",
			expected_status_code:  400,
			expected_request_body: `{"message":"validation error"}`,
		},
		{
			test_name:  "binding error",
			input_body: `{"email": "123@withoutasecondthought.com", "password":""}`,
			input_user: mementor_back.Auth{
				Email:    "test",
				Password: "",
			},
			mock_func:             func(s *mock_service.MockAuthorization, auth mementor_back.Auth) {},
			content_type:          "application/xml",
			expected_status_code:  400,
			expected_request_body: `{"message":"binding error"}`,
		},
	}

	for _, test := range data {
		t.Run(test.test_name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			auth := mock_service.NewMockAuthorization(ctrl)
			test.mock_func(auth, test.input_user)

			services := &service.Service{Authorization: auth}
			handler := NewHandler(services)

			r := echo.New()
			r.POST("/sign-up", handler.signUp)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up", bytes.NewBufferString(test.input_body))
			req.Header.Add("Content-Type", test.content_type)

			r.ServeHTTP(w, req)

			assert.Equal(t, test.expected_status_code, w.Code)
			assert.Equal(t, test.expected_request_body+"\n", w.Body.String())
		})
	}
}

func TestHandlerSignIn(t *testing.T) {
	data := []struct {
		test_name             string
		input_body            string
		input_user            mementor_back.Auth
		mock_func             func(s *mock_service.MockAuthorization, auth mementor_back.Auth)
		content_type          string
		expected_status_code  int
		expected_request_body string
	}{{
		test_name:  "good",
		input_body: `{"email":"123@withoutasecondthought.com", "password":"passed"}`,
		input_user: mementor_back.Auth{
			Email:    "123@withoutasecondthought.com",
			Password: "passed",
		},
		mock_func: func(s *mock_service.MockAuthorization, auth mementor_back.Auth) {
			ctx := context.Background()
			s.EXPECT().SignIn(ctx, auth).Return("1", nil)
		},
		content_type:          "application/json",
		expected_status_code:  200,
		expected_request_body: `{"token":"1"}`,
	},
		{
			test_name:  "empty fields",
			input_body: `{"email": "123@withoutasecondthought.com", "password":""}`,
			input_user: mementor_back.Auth{
				Email:    "test",
				Password: "",
			},
			mock_func:             func(s *mock_service.MockAuthorization, auth mementor_back.Auth) {},
			content_type:          "application/json",
			expected_status_code:  400,
			expected_request_body: `{"message":"validation error"}`,
		},
		{
			test_name:  "email error",
			input_body: `{"email": "@withoutasecondthought", "password":"123321"}`,
			input_user: mementor_back.Auth{
				Email:    "test",
				Password: "",
			},
			mock_func:             func(s *mock_service.MockAuthorization, auth mementor_back.Auth) {},
			content_type:          "application/json",
			expected_status_code:  400,
			expected_request_body: `{"message":"validation error"}`,
		},
		{
			test_name:  "binding error",
			input_body: `{"email": "123@withoutasecondthought.com", "password":""}`,
			input_user: mementor_back.Auth{
				Email:    "test",
				Password: "",
			},
			mock_func:             func(s *mock_service.MockAuthorization, auth mementor_back.Auth) {},
			content_type:          "application/xml",
			expected_status_code:  400,
			expected_request_body: `{"message":"binding error"}`,
		},
	}

	for _, test := range data {
		t.Run(test.test_name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			auth := mock_service.NewMockAuthorization(ctrl)
			test.mock_func(auth, test.input_user)

			services := &service.Service{Authorization: auth}
			handler := NewHandler(services)

			r := echo.New()
			r.POST("/sign-in", handler.signIn)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-in", bytes.NewBufferString(test.input_body))
			req.Header.Add("Content-Type", test.content_type)

			r.ServeHTTP(w, req)

			assert.Equal(t, test.expected_status_code, w.Code)
			assert.Equal(t, test.expected_request_body+"\n", w.Body.String())
		})
	}
}
