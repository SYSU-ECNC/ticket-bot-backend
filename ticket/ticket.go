package ticket

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type Ticket struct {
	ID           string
	TicketStatus string `form:"TicketStatus"  json:"TicketStatus"`
	TicketFrom   string `form:"TicketFrom"  json:"TicketFrom"`
	TicketLabel  string `form:"TicketLabel"  json:"TicketLabel"`
	Creator      string `form:"Creator"  json:"Creator"`
	LastUpdate   string
}

// NewTicket 新建工单()
func NewTicket(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"状态": "你新建了一个工单",
	})
}

// BuildTicket 收集工单信息
func BuildTicket(c *gin.Context) {
	var ticket Ticket
	now := time.Now()
	NowTime := now.Year()*100000000 + int(now.Month())*1000000 + now.Day()*10000 + now.Hour()*100 + now.Minute()
	//TicketStatus := c.DefaultPostForm("状态0", "未解决")
	//TicketFrom := c.DefaultPostForm("来源0", "电话")
	//TicketLabel := c.DefaultPostForm("标签0", "")
	//Creator := c.PostForm("创建者") // 必填，后续如果有用户信息可以设为默认
	if err := c.ShouldBind(&ticket); err == nil {
		ticket.ID = strconv.Itoa(NowTime) // 用当时时间作为ID
		ticket.LastUpdate = strconv.Itoa(NowTime)
		c.JSON(http.StatusCreated, gin.H{
			"ID":   ticket.ID,
			"状态":   ticket.TicketStatus,
			"来源":   ticket.TicketFrom,
			"标签":   ticket.TicketLabel,
			"创建者":  ticket.Creator,
			"更新时间": ticket.LastUpdate,
		})
	} else {
		c.JSON(http.StatusBadGateway, gin.H{
			"错误": err.Error(),
		})
	}

}

// ShowTicket 显示工单信息（未实现）
func ShowTicket(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"显示": "显示工单",
	})
}

// UpdateTicket 更新工单(未实现）
func UpdateTicket(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{
		"更新": "更新工单",
	})
}

// DeleteTicket 删除工单（未实现）
func DeleteTicket(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{
		"删除": "删除工单",
	})
}
