package test

import (
	internal "VentureCookie1/internal/http"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func MockJson(method string, c *gin.Context, content interface{}) {
	c.Request.Method = method
	c.Request.Header.Set("Content-Type", "application/json")

	jsonbytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))
}

func TestPostUser(t *testing.T) {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	ctx.Request = &http.Request{
		Header: make(http.Header),
	}

	MockJson("POST", ctx, map[string]interface{}{"visited": "https://test.com"})

	internal.PostUser(ctx)
	assert.EqualValues(t, http.StatusOK, w.Code)
}

func TestPutUser(t *testing.T) {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	ctx.Request = &http.Request{
		Header: make(http.Header),
	}

	MockJson("PUT", ctx, map[string]interface{}{"visited": "https://test.com", "userid": "617c4c51334b3e8ca548b62c"})

	internal.PostUser(ctx)
	assert.EqualValues(t, http.StatusOK, w.Code)
}
