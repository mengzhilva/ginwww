package main

import (
	"ginwww/router"
	"ginwww/crontab"
)

func main() {
	crontab.Cron()
	router.InitRouter()
}

