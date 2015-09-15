package main

import (
	"fmt"
	"github.com/toshi3221/theta_v2"
	"os"
)

func main() {

	host := "http://192.168.1.1"
	if len(os.Args) > 1 {
		host = os.Args[1]
	}

	client, _ := theta_v2.NewClient(host)
	res, _ := client.State()

	fmt.Println("fingerprint:", *res.Fingerprint)
	fmt.Println("state:")
	fmt.Println("  sessionId:", *res.State.SessionId)
	fmt.Println("  battelyLevel:", *res.State.BatteryLevel)

}
