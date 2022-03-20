package router

import (
	"github.com/gin-gonic/gin"
	"ticket-bot/ticket"
)

func SetupRouters() *gin.Engine {
	r := gin.Default()

	r.GET("/new", ticket.NewTicket)
	r.POST("/build", ticket.BuildTicket)
	r.GET("/show", ticket.ShowTicket)
	r.POST("/update", ticket.UpdateTicket)
	r.DELETE("/delete", ticket.DeleteTicket)
	return r
}
