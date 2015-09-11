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
	info, _ := client.Info()

	fmt.Println("manufacturer:", *info.Manufacturer)
	fmt.Println("model:", *info.Model)
	fmt.Println("serialNumber:", *info.SerialNumber)
	fmt.Println("firmwareVersion:", *info.FirmwareVersion)

}
