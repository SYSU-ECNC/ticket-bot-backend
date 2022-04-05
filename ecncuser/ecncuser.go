package ecncuser

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	resp, _ := http.Get("https://accounts.ecnc.link/kratos/sessions/whoami")
	if resp.StatusCode == 401 {
		//c.Redirect(302, "https://accounts.ecnc.link/kratos/self-service/login/browser?return_to=https://localhost.staging.ecnc.link:3000/show")
		//调试时该链接错误
		c.Redirect(302, "https://accounts.ecnc.link")
	} else {
		
		c.JSON(200, gin.H{
			"login": "login in",
			"resp":  resp.StatusCode,
		})
	}
	defer resp.Body.Close()
}
