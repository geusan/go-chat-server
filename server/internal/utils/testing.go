package testing_utils

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func NewTestHttp(t *testing.T, ctx context.Context, method string, url string, body io.Reader) (c echo.Context, req *http.Request, res *httptest.ResponseRecorder) {
	e := echo.New()
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	assert.NoError(t, err)
	res = httptest.NewRecorder()
	c = e.NewContext(req, res)
	return c, req, res
}
