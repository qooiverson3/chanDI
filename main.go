package main

import (
	"chanLoader/pkg/domain"
	"chanLoader/pkg/handler"
	"chanLoader/pkg/service"

	"github.com/gin-gonic/gin"
)

const URL = "mongodb://root:example@localhost:27017"

func main() {
	// for {
	// 	v, ok := <-ch
	// 	if !ok {
	// 		break
	// 	}
	// 	re = append(re, v)
	// }

	router := gin.Default()

	data1 := domain.Info{
		Data: []int{1, 2, 3, 4, 5},
		Name: "Data-1",
	}
	data2 := domain.Info{
		Data: []int{6, 7, 8, 9, 10},
		Name: "Data-2",
	}
	data3 := domain.Info{
		Data: []int{6, 7, 8, 9, 10},
		Name: "Data-3",
	}

	list := []domain.Info{
		data1, data2, data3,
	}

	chanService := make(chan domain.Service, len(list))
	for i := 0; i < len(list); i++ {
		go func(n int) {
			s := service.New(list[n])
			chanService <- s

			s.Exporter(list[n].Name)
		}(i)
	}

	h := handler.New(chanService, len(list))
	h.Route(router)

	router.Run(":8080")
}
