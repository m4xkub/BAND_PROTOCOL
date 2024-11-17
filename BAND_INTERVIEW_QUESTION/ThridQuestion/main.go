package main

import (
	"ThridQuestion/internal/config"
	"ThridQuestion/internal/controller"
	"fmt"
	"time"

	"ThridQuestion/internal/service"

	"github.com/go-co-op/gocron"
)

func main() {
	cron := gocron.NewScheduler(time.UTC)

	cron.Every(config.MonitorInterval).Do(func() {
		currentTime := time.Now().Format("2006-01-02 15:04:05")

		isChange := true

		fmt.Printf("[%s] ---------------------------------------------------------------------------\n", currentTime)
		for i := 0; i < len(service.TransactionMemory); i++ {
			service.Monitor(&service.TransactionMemory[i], &isChange)
		}

		if isChange {
			fmt.Println("All transactions' status are in finalization state")
		}
		fmt.Printf("-------------------------------------------------------------------------------------------------\n")

	})

	cron.StartAsync()

	router := controller.GetRootController()

	router.Run(":8080")
}
