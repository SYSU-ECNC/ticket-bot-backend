package router

import (
	"github.com/gin-gonic/gin"
	"ticket-bot/ecncuser"
	"ticket-bot/ticket"
)

func SetupRouters() *gin.Engine {
	r := gin.Default()
	r.GET("/login", ecncuser.Login)

	//r.GET("/new", ticket.NewTicket)

	r.POST("/tickets", ticket.BuildTicket)
	r.GET("/tickets", ticket.ShowTickets)
	r.GET("/tickets/:ID", ticket.ShowTicket)
	r.PATCH("/tickets/:ID", ticket.PatchTicket)
	r.DELETE("/tickets/:ID", ticket.DeleteTicket)
	return r
}
