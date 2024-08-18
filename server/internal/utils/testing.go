package testing_utils

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func NewTestHttp[T interface{}](t *testing.T, ctx context.Context, method string, url string, body T) (c echo.Context, req *http.Request, res *httptest.ResponseRecorder) {
	e := echo.New()
	buf, err := json.Marshal(body)
	assert.NoError(t, err)
	req, err = http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(buf))
	assert.NoError(t, err)
	res = httptest.NewRecorder()
	c = e.NewContext(req, res)
	req.Header.Set("Content-Type", "application/json")
	return c, req, res
}
