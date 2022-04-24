package ecncuser

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
)

var larkUnionId string
var netId string

func Login(c *gin.Context) {
	resp, _ := http.Get("https://accounts.ecnc.link/kratos/sessions/whoami")
	if resp.StatusCode == 401 {
		c.Redirect(302, "https://accounts.ecnc.link")
	} else {
		body, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(body)
		uidString := bodyString[strings.Index(bodyString, "lark_union_id")+16:]
		larkUnionId = uidString[:strings.Index(uidString, "\"")]
		netString := bodyString[strings.Index(bodyString, "netid")+8:]
		netId = netString[:strings.Index(netString, "\"")]

		c.JSON(200, gin.H{
			"login": "login in",
			"resp":  resp.StatusCode,
		})

	}
	defer resp.Body.Close()
}

func GetUnionID() string {
	return larkUnionId
}

func GetNetID() string {
	return netId
}
