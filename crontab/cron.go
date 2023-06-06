package crontab

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

// Cron 定时任务
func Cron() {
	myCron := cron.New()
	//计划任务，每分钟
	addFunc, cronErr := myCron.AddFunc("@every 1m", func() {
		fmt.Println("每分钟执行一次")
	})
	if cronErr != nil {
		//myCron.Stop()
		fmt.Println(addFunc, cronErr)
	}
	addFunc1, cronErr := myCron.AddFunc("@every 1m", func() {
		fmt.Println("第二个每分钟执行一次")
	})
	if cronErr != nil {
		//myCron.Stop()
		fmt.Println(addFunc1, cronErr)
	}
	myCron.Start()
}
