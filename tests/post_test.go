package tests

import (
	"database/sql/driver"
	"echo-demo-project/requests"
	"echo-demo-project/server"
	"echo-demo-project/server/handlers"
	"echo-demo-project/services/token"
	"echo-demo-project/tests/helpers"
	"net/http"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var postId, _ = uuid.NewV7()

var postIdNotExists, _ = uuid.NewV7()

func TestWalkPostsCrud(t *testing.T) {
	dbMock, sqlMock, err := sqlmock.New()
	if err != nil {
		panic(err.Error())
	}

	requestCreate := helpers.Request{
		Method: http.MethodPost,
		Url:    "/posts",
	}
	requestGet := helpers.Request{
		Method: http.MethodGet,
		Url:    "/posts",
	}
	requestUpdate := helpers.Request{
		Method: http.MethodPut,
		Url:    "/posts/" + (postId).String(),
		PathParam: &helpers.PathParam{
			Name:  "uuid",
			Value: (postId).String(),
		},
	}
	requestDelete := helpers.Request{
		Method: http.MethodDelete,
		Url:    "/posts/" + (postId).String(),
		PathParam: &helpers.PathParam{
			Name:  "id",
			Value: (postId).String(),
		},
	}
	handlerFuncCreate := func(s *server.Server, c echo.Context) error {
		return handlers.NewPostHandlers(s).CreatePost(c)
	}
	handlerFuncGet := func(s *server.Server, c echo.Context) error {
		return handlers.NewPostHandlers(s).GetPosts(c)
	}
	handlerFuncUpdate := func(s *server.Server, c echo.Context) error {
		return handlers.NewPostHandlers(s).UpdatePost(c)
	}
	handlerFuncDelete := func(s *server.Server, c echo.Context) error {
		return handlers.NewPostHandlers(s).DeletePost(c)
	}

	claims := &token.JwtCustomClaims{
		Name: "user",
		UUID: helpers.UserId,
	}
	validToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	commonMock := &helpers.QueryMock{
		Query:    "SELECT * FROM `posts` WHERE uuid = ? AND `posts`.`deleted_at` IS NULL",
		QueryArg: []driver.Value{helpers.UserId},
		Reply: helpers.MockReply{
			Columns: []string{"uuid", "title", "content", "username"},
			Rows: [][]driver.Value{
				{helpers.UserId, "title", "content", "Username"},
			},
		},
	}

	cases := []helpers.TestCase{
		{
			"Create post success",
			requestCreate,
			requests.CreatePostRequest{
				BasicPost: requests.BasicPost{
					Title:   "title",
					Content: "content",
				},
			},
			handlerFuncCreate,
			[]*helpers.QueryMock{&helpers.SelectVersionMock},
			helpers.ExpectedResponse{
				StatusCode: 201,
				BodyPart:   "Post successfully created",
			},
		},
		{
			"Create post with empty title",
			requestCreate,
			requests.CreatePostRequest{
				BasicPost: requests.BasicPost{
					Title:   "",
					Content: "content",
				},
			},
			handlerFuncCreate,
			[]*helpers.QueryMock{&helpers.SelectVersionMock},
			helpers.ExpectedResponse{
				StatusCode: 400,
				BodyPart:   "Required fields are empty",
			},
		},
		{
			"Get posts success",
			requestGet,
			"",
			handlerFuncGet,
			[]*helpers.QueryMock{
				&helpers.SelectVersionMock,
				{
					Query: "SELECT * FROM `posts` WHERE ",
					Reply: helpers.MockReply{
						Columns: []string{"uuid", "title", "content", "username"},
						Rows: [][]driver.Value{
							{helpers.UserId, "title", "content", "Username"},
						},
					},
				},
			},
			helpers.ExpectedResponse{
				StatusCode: 200,
				BodyPart:   "[{\"title\":\"title\",\"content\":\"content\",\"username\":\"\",\"uuid\":1}]",
			},
		},
		{
			"Update post success",
			requestUpdate,
			requests.UpdatePostRequest{
				BasicPost: requests.BasicPost{
					Title:   "new title",
					Content: "new content",
				},
			},
			handlerFuncUpdate,
			[]*helpers.QueryMock{&helpers.SelectVersionMock, commonMock},
			helpers.ExpectedResponse{
				StatusCode: 200,
				BodyPart:   "Post successfully updated",
			},
		},
		{
			"Update post with empty title",
			requestUpdate,
			requests.UpdatePostRequest{
				BasicPost: requests.BasicPost{
					Title:   "",
					Content: "new content",
				},
			},
			handlerFuncUpdate,
			[]*helpers.QueryMock{&helpers.SelectVersionMock},
			helpers.ExpectedResponse{
				StatusCode: 400,
				BodyPart:   "Required fields are empty",
			},
		},
		{
			"Update non-existent post",
			helpers.Request{
				Method: http.MethodPut,
				Url:    "/posts/" + (postIdNotExists).String(),
				PathParam: &helpers.PathParam{
					Name:  "uuid",
					Value: (postIdNotExists).String(),
				},
			},
			requests.UpdatePostRequest{
				BasicPost: requests.BasicPost{
					Title:   "new title",
					Content: "new content",
				},
			},
			handlerFuncUpdate,
			[]*helpers.QueryMock{&helpers.SelectVersionMock},
			helpers.ExpectedResponse{
				StatusCode: 404,
				BodyPart:   "Post not found",
			},
		},
		{
			"Delete post success",
			requestDelete,
			"",
			handlerFuncDelete,
			[]*helpers.QueryMock{&helpers.SelectVersionMock, commonMock},
			helpers.ExpectedResponse{
				StatusCode: 204,
				BodyPart:   "Post deleted successfully",
			},
		},
		{
			"Delete non-existent post",
			helpers.Request{
				Method: http.MethodDelete,
				Url:    "/posts/" + (postIdNotExists).String(),
				PathParam: &helpers.PathParam{
					Name:  "uuid",
					Value: (postIdNotExists).String(),
				},
			},
			"",
			handlerFuncDelete,
			[]*helpers.QueryMock{&helpers.SelectVersionMock, commonMock},
			helpers.ExpectedResponse{
				StatusCode: 404,
				BodyPart:   "Post not found",
			},
		},
	}

	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {
			helpers.PrepareDatabaseQueryMocks(test, sqlMock)
			db := helpers.InitGorm(dbMock)
			s := helpers.NewServer(db)

			c, recorder := helpers.PrepareContextFromTestCase(s, test)
			c.Set("user", validToken)

			if assert.NoError(t, test.HandlerFunc(s, c)) {
				assert.Contains(t, recorder.Body.String(), test.Expected.BodyPart)
				assert.Equal(t, test.Expected.StatusCode, recorder.Code)
			}
		})
	}
}
