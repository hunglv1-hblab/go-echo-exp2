package helpers

import (
	"echo-demo-project/server"
	"encoding/json"
	"github.com/labstack/echo/v4"
	mocket "github.com/selvatico/go-mocket"
	"net/http/httptest"
	"strings"
)

const UserId = 1

type MockReply []map[string]interface{}

type Request struct {
	Method string
	Url    string
}

type QueryMock struct {
	Query string
	Reply MockReply
}

type ExpectedResponse struct {
	StatusCode int
	BodyPart   string
}

type TestCase struct{
	TestName    string
	Request     Request
	RequestBody interface{}
	HandlerFunc func(s *server.Server, c echo.Context) error
	QueryMock   *QueryMock
	Expected    ExpectedResponse
}

func PrepareContextFromTestCase(s *server.Server, test TestCase) (c echo.Context, recorder *httptest.ResponseRecorder) {
	requestJson, _ := json.Marshal(test.RequestBody)
	request := httptest.NewRequest(test.Request.Method, test.Request.Url, strings.NewReader(string(requestJson)))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder = httptest.NewRecorder()
	c = s.Echo.NewContext(request, recorder)

	if test.QueryMock != nil {
		mocket.Catcher.Reset().NewMock().WithQuery(test.QueryMock.Query).WithReply(test.QueryMock.Reply)
	}

	return
}