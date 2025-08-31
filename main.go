package main

import (
	"fmt"
	"github.com/amarnathcjd/gogram/telegram"
	"log"
)

func main() {
	var (
		apiID   int32
		apiHash string
	)
	fmt.Print("Enter an API ID: ")
	if _, err := fmt.Scan(&apiID); err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Print("Enter an API HASH: ")
	if _, err := fmt.Scan(&apiHash); err != nil {
		log.Fatalf(err.Error())
	}
	ubot, err := telegram.NewClient(telegram.ClientConfig{
		AppID:         apiID,
		AppHash:       apiHash,
		LogLevel:      telegram.LogWarn,
		NoUpdates:     true,
		DisableCache:  true,
		MemorySession: true,
	})
	ubot.Log.NoColor()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println(ubot.Me().FirstName, "has been started!")
	fmt.Println(ubot.ExportSession())
	_ = ubot.Stop()
	log.Println("Bye!")
}
