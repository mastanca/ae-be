package handlers

import (
	"net/http"
	"testing"

	"github.com/mastanca/go-api-template/test"
	"github.com/stretchr/testify/assert"
)

func TestPingHandlerImpl_Handle(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		handler := NewPingHandlerImpl()
		router := test.Router("/api/ping", handler.Handle, http.MethodGet)

		response := test.MakeRequest(router, http.MethodGet, "/api/ping", nil)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, "pong", response.Body.String())
	})
}
