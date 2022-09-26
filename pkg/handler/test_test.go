package handler_test

import (
	"chanLoader/pkg/domain"
	"chanLoader/pkg/handler"
	"chanLoader/pkg/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetStreams(t *testing.T) {
	fakeData := domain.Info{
		Name: "data-1",
		Data: []int{1, 2, 3},
	}
	fakeAmount := 2
	fakeCh := make(chan domain.Service, fakeAmount)
	r := gin.Default()

	for i := 0; i < fakeAmount; i++ {
		s := service.New(fakeData)
		fakeCh <- s
	}

	h := handler.New(fakeCh, fakeAmount)
	h.Route(r)

	t.Run("should return http 200 code", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/info", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})
}
