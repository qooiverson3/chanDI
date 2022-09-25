package service_test

import (
	"chanLoader/pkg/domain"
	"chanLoader/pkg/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetData(t *testing.T) {
	fake := domain.Info{
		Data: []int{
			1, 2, 3, 4, 5,
		},
		Name: "Data-1",
	}
	t.Run("should get data-1 struct", func(t *testing.T) {
		s := service.New(fake)
		data := s.GetData()
		assert.Equal(t, fake.Data, data.Data)
		assert.Equal(t, fake.Name, data.Name)
	})
}
