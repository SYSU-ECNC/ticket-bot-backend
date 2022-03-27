package ticket

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type Ticket struct {
	ID           string
	TicketStatus string `json:"ticketStatus"`
	TicketFrom   string `json:"ticketFrom"`
	TicketLabel  string `json:"ticketLabel"`
	Creator      string `json:"creator"`
	LastUpdate   string
	BMC          string
}

//// NewTicket 新建工单()
//func NewTicket(c *gin.Context) {
//	c.JSON(http.StatusOK, gin.H{
//		"状态": "你新建了一个工单",
//	})
//}

// BuildTicket 收集工单信息
func BuildTicket(c *gin.Context) {
	var ticket Ticket
	now := time.Now()
	NowTime := strconv.Itoa(now.Year()) + "-" + strconv.Itoa(int(now.Month())) + "-" + strconv.Itoa(now.Day()) + "-" + strconv.Itoa(now.Hour()) + ":" + strconv.Itoa(now.Minute())
	ticket.TicketStatus = "未解决"
	ticket.TicketFrom = "电话"
	//TicketLabel := c.DefaultPostForm("标签0", "")
	//Creator := c.PostForm("创建者") // 必填，后续如果有用户信息可以设为默认
	if err := c.ShouldBind(&ticket); err == nil {
		ticket.ID = NowTime // 用当时时间作为ID
		ticket.LastUpdate = NowTime
		c.JSON(http.StatusCreated, gin.H{
			"ID":           ticket.ID,
			"ticketStatus": ticket.TicketStatus,
			"ticketFrom":   ticket.TicketFrom,
			"ticketLabel":  ticket.TicketLabel,
			"creator":      ticket.Creator,
			"lastUpdate":   ticket.LastUpdate,
			"BMC":          ticket.BMC,
		})
		//ChatID := ticket.ID
		//robot.NewChat(&ChatID)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

}

// ShowTickets 显示工单信息（未实现）
func ShowTickets(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"show": "显示全部工单",
	})
}

// ShowTicket 显示工单（未实现）
func ShowTicket(c *gin.Context) {
	ID := c.Param("ID")
	c.JSON(http.StatusOK, gin.H{
		"show": "显示" + ID + "工单",
	})
}

// PatchTicket 更新工单(未实现）
func PatchTicket(c *gin.Context) {
	ID := c.Param("ID")
	c.JSON(http.StatusAccepted, gin.H{
		"patch": "更新" + ID + "工单",
	})
}

// DeleteTicket 删除工单（未实现）
func DeleteTicket(c *gin.Context) {
	ID := c.Param("ID")
	c.JSON(http.StatusAccepted, gin.H{
		"delete": "删除" + ID + "工单",
	})
}
