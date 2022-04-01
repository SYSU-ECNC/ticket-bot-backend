package router

import (
	"github.com/gin-gonic/gin"
	"ticket-bot-backend/ticket"
)

func SetupRouters() *gin.Engine {
	r := gin.Default()

	r.POST("/tickets/create", ticket.CreateTicket) //crteate a new ticket
	r.GET("/tickets/show", ticket.ShowAllTickets)  //list all tickets
	r.GET("/tickets/:id", ticket.ShowTicket)       //list the specified ticket
	r.GET("/tickets", ticket.GetUnfishedTicket)    //list unfished tickets
	r.PATCH("/tickets/:id", ticket.UpdateTicket)   //update some information of the tickets
	r.DELETE("/tickets/:id", ticket.DeleteTicket)  //delete a ticket
	return r
}
