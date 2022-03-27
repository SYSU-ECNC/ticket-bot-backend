package ecncuser

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	resp, _ := http.Get("https://accounts.ecnc.link/kratos/sessions/whoami")
	ok := "ok"
	if resp.StatusCode == 401 {
		ok = "not ok"
		c.Redirect(302, "https://accounts.ecnc.link/kratos/self-service/login/browser?return_to=https://localhost.staging.ecnc.link:3000/show")
		resp, _ = http.Get("https://accounts.ecnc.link/kratos/sessions/whoami")
		c.JSON(200, gin.H{
			"no": resp.StatusCode,
		})
	} else {
		c.JSON(200, gin.H{
			"login": "login",
			"resp":  resp.StatusCode,
			"ok":    ok,
		})
	}
	//fmt.Println(resp)
	defer resp.Body.Close()
}
