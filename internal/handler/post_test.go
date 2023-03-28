package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	mock "go-boilerplate/internal/handler/mock"
	"go-boilerplate/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type PostHandlerSuite struct {
	mUsecase       *mock.MockPostUsecase
	responseWriter *httptest.ResponseRecorder
	router         *gin.Engine
}

func initPostTestSuite(t *testing.T) PostHandlerSuite {
	ctrl := gomock.NewController(t)

	return PostHandlerSuite{
		mUsecase: mock.NewMockPostUsecase(ctrl),
		router:   gin.Default(),
	}
}

func (suite *PostHandlerSuite) makeHTTPCall(method, url, reqBody string) {
	suite.responseWriter = httptest.NewRecorder()
	request, _ := http.NewRequest(method, url, strings.NewReader(reqBody))
	request.Header.Add("Content-Type", "application/json")

	suite.router.ServeHTTP(suite.responseWriter, request)
}

func TestHandler_Hello(t *testing.T) {
	suite := initPostTestSuite(t)

	postHandler := NewPostHandler(suite.mUsecase)
	suite.router.GET("/", postHandler.Hello)

	post := model.Post{
		Title:   "hello",
		Content: "hello world",
	}

	t.Run("success", func(t *testing.T) {
		suite.makeHTTPCall(http.MethodGet, "/", "")

		resBytes, _ := json.MarshalIndent(newResponse(post), "", "    ")
		resBody := string(resBytes)

		assert.Equal(t, http.StatusOK, suite.responseWriter.Code)
		assert.Equal(t, resBody, suite.responseWriter.Body.String())
	})
}

func TestHandler_GetPosts(t *testing.T) {
	suite := initPostTestSuite(t)

	postHandler := NewPostHandler(suite.mUsecase)
	suite.router.GET("/posts", postHandler.GetPosts)

	posts := []model.Post{
		{
			Title:   "hello",
			Content: "hello world",
		},
	}

	t.Run("internal server error", func(t *testing.T) {
		suite.mUsecase.EXPECT().GetPosts(gomock.Any()).Return(nil, errors.New("error"))
		suite.makeHTTPCall(http.MethodGet, "/posts", "")

		assert.Equal(t, http.StatusInternalServerError, suite.responseWriter.Code)
	})

	t.Run("success", func(t *testing.T) {
		suite.mUsecase.EXPECT().GetPosts(gomock.Any()).Return(posts, nil)
		suite.makeHTTPCall(http.MethodGet, "/posts", "")

		resBytes, _ := json.MarshalIndent(newResponse(posts), "", "    ")
		resBody := string(resBytes)

		assert.Equal(t, http.StatusOK, suite.responseWriter.Code)
		assert.Equal(t, resBody, suite.responseWriter.Body.String())
	})
}

func TestHandler_CreatePost(t *testing.T) {
	suite := initPostTestSuite(t)

	postHandler := NewPostHandler(suite.mUsecase)
	suite.router.POST("/posts", postHandler.CreatePost)

	req := model.Post{
		Title:   "new post",
		Content: "post",
	}
	reqBytes, _ := json.Marshal(req)
	reqBody := string(reqBytes)

	t.Run("success", func(t *testing.T) {
		suite.mUsecase.EXPECT().CreatePost(gomock.Any(), req).Return(nil)
		suite.makeHTTPCall(http.MethodPost, "/posts", reqBody)

		assert.Equal(t, http.StatusCreated, suite.responseWriter.Code)
	})

	t.Run("invalid body request", func(t *testing.T) {
		suite.makeHTTPCall(http.MethodPost, "/posts", "{")
		assert.Equal(t, http.StatusBadRequest, suite.responseWriter.Code)
	})

	t.Run("internal server error", func(t *testing.T) {
		suite.mUsecase.EXPECT().CreatePost(gomock.Any(), req).Return(errors.New("error"))
		suite.makeHTTPCall(http.MethodPost, "/posts", reqBody)

		assert.Equal(t, http.StatusInternalServerError, suite.responseWriter.Code)
	})
}
