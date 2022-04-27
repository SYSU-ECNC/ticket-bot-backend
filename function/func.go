package function

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"ticket-bot-backend/ecncuser"
	"ticket-bot-backend/robot"
	"ticket-bot-backend/ticket"
	"time"
)

//CreateTicket 收集工单信息
func CreateTicket(c *gin.Context) {

	db, err := ticket.OpenDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var thisticket ticket.Ticket
	err = c.BindJSON(&thisticket)
	if err != nil {
		log.Println("BindJSON Failed", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	thisticket.ID = uint(time.Now().Unix())
	db.AutoMigrate(&ticket.Ticket{})
	chatId := robot.NewChat(thisticket, ecncuser.GetUnionID(), ecncuser.GetNetID())
	thisticket.RelatedInf.FeishuGroup = chatId

	result := db.Create(&thisticket)
	if result.Error != nil {
		log.Println("Create Error", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"ticket": thisticket,
		})
	}
}

func Getpost(c *gin.Context) {
	postJson := struct {
		Event struct {
			Sender struct {
				Senderid struct {
					Unionid string `json:"union_id"`
					Userid  string `json:"user_id"`
				} `json:"sender_id"`
			} `json:"sender"`
			Mess struct {
				Messageid  string `json:"message_id"`
				Chatid     string `json:"chat_id"`
				Context    string `json:"context"`
				Createtime string `json:"create_time"`
			} `json:"message"`
		} `json:"event"`
	}{}
	err := c.BindJSON(&postJson)
	if err != nil {
		fmt.Printf("bind error:%v\n", err)
	} else {
		db, _ := ticket.OpenDB()
		var oldticket ticket.Ticket
		db.First(&oldticket, postJson.Event.Mess.Chatid)

		var updateinfo ticket.Ticket
		updateinfo = oldticket
		var newchatrecord ticket.Messagerecord
		newchatrecord.Createtime = postJson.Event.Mess.Createtime
		newchatrecord.Unionid = postJson.Event.Sender.Senderid.Unionid
		newchatrecord.Chatid = postJson.Event.Mess.Chatid
		newchatrecord.Context = postJson.Event.Mess.Context
		newchatrecord.Messageid = postJson.Event.Mess.Messageid
		newchatrecord.Userid = postJson.Event.Sender.Senderid.Userid
		newrecord := append(updateinfo.RelatedInf.Chatrecord, newchatrecord)
		updateinfo.RelatedInf.Chatrecord = newrecord
		db.Model(&oldticket).Updates(updateinfo)

	}

	c.JSON(200, gin.H{})
}
