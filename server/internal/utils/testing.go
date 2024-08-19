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

func buildBodyReader(rawBody interface{}) ([]byte, error) {
	if t, ok := rawBody.(string); ok {
		return []byte(t), nil
	}
	if rawBody == nil {
		return []byte{}, nil
	}
	buff, err := json.Marshal(rawBody)
	return []byte(buff), err
}

func NewTestHttp(t *testing.T, ctx context.Context, method string, url string, rawBody interface{}) (c echo.Context, req *http.Request, res *httptest.ResponseRecorder) {
	e := echo.New()
	body, err := buildBodyReader(rawBody)
	assert.NoError(t, err)
	req, err = http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(body))
	assert.NoError(t, err)
	res = httptest.NewRecorder()
	c = e.NewContext(req, res)
	req.Header.Set("Content-Type", "application/json")
	return c, req, res
}
