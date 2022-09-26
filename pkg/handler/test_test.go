package handler_test

import (
	"chanLoader/pkg/domain"
	"chanLoader/pkg/handler"
	"chanLoader/pkg/service"
	"chanLoader/pkg/storage/fake"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetStreams(t *testing.T) {

	fakeCh := make(chan domain.Service, fake.FakeAmount)
	r := gin.Default()

	for i := 0; i < fake.FakeAmount; i++ {
		s := service.New(fake.FakeDataList[i])
		fakeCh <- s
	}

	h := handler.New(fakeCh, fake.FakeAmount)
	h.Route(r)

	t.Run("should return http 200 code", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/info", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, `{"data_list":[[1,2,3],[11,22,33]],"name_list":["fake-data-1","fake-data-2"]}`, w.Body.String())
	})
}

func TestGetOneStreamInfo(t *testing.T) {

	fakeCh := make(chan domain.Service, 1)

	s := service.New(fake.FakeData1)
	fakeCh <- s

	//ch := <-fakeCh
	r := gin.Default()
	h := handler.New(fakeCh, 1)
	h.Route(r)

	t.Run("should return http 200 code", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/streamInfo/fake-data-1", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		//assert.Equal(t, fmt.Sprintf("{\"info\":%v}", ch.GetData()), w.Body.String())
	})
}
