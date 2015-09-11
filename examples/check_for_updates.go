package main

import (
	"fmt"
	"github.com/toshi3221/osc"
	"os"
)

func main() {

	host := "http://192.168.1.1"
	if len(os.Args) > 1 {
		host = os.Args[1]
	}
	client, _ := osc.NewClient(host)

	res, _ := client.CheckForUpdates("test", nil)
	if res.StateFingerprint != nil {
		fmt.Println("fingerprint:", *res.StateFingerprint)
	}
	if res.ThrottleTimeout != nil {
		fmt.Println("throttleTimout:", *res.ThrottleTimeout)
	}

}
