package main

import (
	// "os"
	"fmt"
	"time"
	"github.com/mruebush/customerpref"
)

type CustomerPref struct {
	Id string
	NotificationPref string
}

// var (
// 	messageBrokerHost  = os.Getenv("MESSAGE_BROKER")
// 	fileServerPort     = os.Getenv("HTTP_FILE_SERVER_PORT")
// 	eventBusServerHost = os.Getenv("HTTP_EVENTBUS_SERVER_HOST")
// 	eventBusServerPort = os.Getenv("HTTP_EVENTBUS_SERVER_PORT")
// )

func main() {
//should get configuration and setup the program through struct into Start
	fmt.Printf("Starting up server and connecting %s\n", time.Now())

	customerpref.StartServer()

}