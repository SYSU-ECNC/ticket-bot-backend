package main

import "ticket-bot-backend/robot"

func main() {
	//r := router.SetupRouters()

	//err := r.RunTLS(":3000", "localhost.staging.ecnc.link.crt", "localhost.staging.ecnc.link.key")
	//if err != nil {
	//	return
	//}

	//r.Run(":3000")

	//resp, _ := http.Get("https://accounts.ecnc.link/kratos/sessions/whoami")
	//if resp.StatusCode == 401 {
	//	http.RedirectHandler("https://accounts.ecnc.link/kratos/self-service/login/browser?return_to=/kratos/sessions/whoami", 302)
	//}
	//fmt.Println(resp.StatusCode)

	testid := "test"
	unionid := "****"
	netid := "*****"
	robot.NewChat(testid, unionid, netid)

}
