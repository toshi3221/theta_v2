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
	info, _ := client.Info()

	fmt.Println("manufacturer:", *info.Manufacturer)
	fmt.Println("model:", *info.Model)
	fmt.Println("serialNumber:", *info.SerialNumber)
	fmt.Println("firmwareVersion:", *info.FirmwareVersion)

}
