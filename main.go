package main

import "ticket-bot/router"

func main() {
	r := router.SetupRouters()

	r.Run(":3000")
}
