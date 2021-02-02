/*
   Package testutil provides test helpers for the controller.
*/
package testutil

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"server/domain"
	"server/middleware/errorhandler"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type UsersResponse struct {
	Code    int                     `json:"code"`
	Message string                  `json:"message"`
	Data    []domain.User           `json:"data"`
	Errors  []errorhandler.APIError `json:"errors"`
}

type UserResponse struct {
	Code    int                     `json:"code"`
	Message string                  `json:"message"`
	Data    domain.User             `json:"data"`
	Errors  []errorhandler.APIError `json:"errors"`
}

func AssertResponse(t *testing.T, router *gin.Engine, req *http.Request, wantCode int, want interface{}) {
	t.Helper()
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, wantCode, w.Code)
	if w.Code != 204 {
		switch want := want.(type) {
		case UsersResponse:
			respBody := UsersResponse{}
			err := json.Unmarshal(w.Body.Bytes(), &respBody)
			assert.NoError(t, err)
			assert.Equal(t, want, respBody)
		case UserResponse:
			respBody := UserResponse{}
			err := json.Unmarshal(w.Body.Bytes(), &respBody)
			assert.NoError(t, err)
			assert.Equal(t, want, respBody)
		}

	}
}
