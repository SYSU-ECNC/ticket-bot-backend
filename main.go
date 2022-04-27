package main

import "ticket-bot-backend/router"

func main() {
	r := router.SetupRouters()

	r.Run(":3000")

}
