package handler

import (
	"chanLoader/pkg/domain"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	chanServices chan domain.Service
	amount       int
	service      domain.Service
}

func New(ctxS chan domain.Service, n int) handler {
	return handler{chanServices: ctxS, amount: n}
}

func (h *handler) Route(e *gin.Engine) {

	e.GET("/metrics", h.Metrics())
	v1 := e.Group("/api/v1")
	v1.GET("/info", h.GetStreams)
	v1.GET("/streamInfo/:stream", h.GetOneStreamInfo)

}

func (h *handler) GetStreams(ctx *gin.Context) {
	nameList := []string{}
	dataList := [][]int{}

	for i := 0; i < h.amount; i++ {
		ch := <-h.chanServices
		nameList = append(nameList, ch.GetData().Name)
		dataList = append(dataList, ch.GetData().Data)

		// keep channel always open
		h.chanServices <- ch
	}

	ctx.JSON(http.StatusOK, gin.H{
		"name_list": nameList,
		"data_list": dataList,
	})
}

func (h *handler) chanSwitch(stream string) {
	for i := 0; i < h.amount; i++ {
		ch := <-h.chanServices
		h.chanServices <- ch
		h.service = ch
		if ch.GetData().Name == stream {
			log.Println("hit!!")
			break
		}
	}
}

func (h *handler) GetOneStreamInfo(c *gin.Context) {
	h.chanSwitch(c.Param("stream"))
	cs := h.service
	c.JSON(http.StatusOK, gin.H{
		"info": cs.GetData(),
	})

}
