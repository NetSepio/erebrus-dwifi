package main

import (
	"github.com/NetSepio/erebrus-dwifi/connections/database"
	wifiAlgorithem "github.com/NetSepio/erebrus-dwifi/wifiHandler/algorithem"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	database.Connect()
}

func main2() {
	wifiAlgorithem.RunWifi()
}
