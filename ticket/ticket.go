package ticket

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Ticket struct {
	gorm.Model
	Status     string        `json:"status"`
	Source     string        `json:"source"`
	Label      string        `json:"label"`
	Creator    string        `json:"creator" gorm:"not null"`
	BMCID      string        `json:"bmcid"`
	RelatedInf TicketRelated `json:"relatedInf" gorm:"embedded"`
}

type Clients struct {
	Name       string `json:"name"`
	NetID      string `json:"netID" gorm:"not null"`
	ID         string `json:"id"`       //职工号/学号
	Category   string `json:"category"` //人员类别
	Department string `json:"department"`
	IDNumber   string `json:"id_number"` //身份证号/护照号
	Mail       string `json:"mail"`
	Phone      string `json:"phone"`
}

type TicketRelated struct {
	ID          uint    `json:"id"`
	Client      Clients `gorm:"embedded" json:"client"`
	FeishuGroup string  `json:"feishuGroup" gorm:"column:feishuGroup"`
	Chatrecord  []Messagerecord
	Summary     string `json:"summary"`
}

type Messagerecord struct {
	Unionid    string `json:"union_id"`
	Userid     string `json:"user_id"`
	Messageid  string `json:"message_id"`
	Chatid     string `json:"chat_id"`
	Context    string `json:"context"`
	Createtime string `json:"create_time"`
}

type SimpleTicket struct {
	ID     uint   `json:"id"`
	Status string `json:"status"`
	Label  string `json:"label"`
}

type DBConfigure struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
	Port     string `json:"port"`
	Sslmode  string `json:"sslmode"`
}

func OpenDB() (*gorm.DB, error) {
	fileName := "DBconfigure.json"
	filePtr, err := os.Open(fileName)
	if err != nil {
		log.Println("Open file failed!", err.Error())
		return nil, err
	}

	var configure DBConfigure
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&configure)
	if err != nil {
		log.Println("Decoder failed", err.Error())
		return nil, err
	}

	dsn := "host=" + configure.Host + " user=" + configure.User + " password=" + configure.Password + " dbname=" + configure.Dbname + " port=" + configure.Port + " sslmode=" + configure.Sslmode
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return db, nil
}

// ShowTicket 显示指定的工单信息
func ShowTicket(c *gin.Context) {
	db, err := OpenDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	var ticket Ticket
	result := db.First(&ticket, id)

	if result.Error != nil {
		log.Println("Retrieve One Error", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"info":  "找不到对应工单",
			"error": result.Error,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"ticket": ticket,
		})
	}

}

//显示所有工单的详细信息
func ShowAllTickets(c *gin.Context) {
	db, err := OpenDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var tickets []Ticket
	result := db.Find(&tickets)
	if result.Error != nil {
		log.Println("Retrieve All Error", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"all_ticket":    tickets,
			"total_records": result.RowsAffected,
		})
	}

}

// GetUnfishedTicket 显示未完成的工单，只显示部分信息
func GetUnfishedTicket(c *gin.Context) {
	db, err := OpenDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	//以下两种方式的结果不同，但不清楚区别是什么
	//var tickets []SimpleTicket
	//result := db.Table("tickets").Select("id", "status", "label").Where("status <> ?", "已解决").Find(&tickets)
	var tickets []Ticket
	result := db.Select("id", "status", "label").Where("status <> ?", "已解决").Find(&tickets)

	if result.Error != nil {
		log.Println("Retrieve Error", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"tickets":       tickets,
			"total_records": result.RowsAffected,
		})
	}

}

// UpdateTicket
func UpdateTicket(c *gin.Context) {
	db, err := OpenDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	var oldticket Ticket
	db.First(&oldticket, id)

	var updateinfo map[string]interface{}
	err = c.BindJSON(&updateinfo)
	if err != nil {
		log.Println("Update Error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result := db.Model(&oldticket).Updates(updateinfo)
	if result.Error != nil {
		log.Println("Update Error", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"tickets":       oldticket,
			"total_records": result.RowsAffected,
		})
	}

}

// DeleteTicket
func DeleteTicket(c *gin.Context) {
	db, err := OpenDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	var ticket Ticket
	result := db.First(&ticket, id)
	if result.Error != nil {
		log.Println("Delete Error", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"info":  "找不到对应工单",
			"error": err,
		})
		return
	}

	result = db.Delete(&ticket)
	if result.Error != nil {
		log.Println("Delete Error", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{})
	}
}
