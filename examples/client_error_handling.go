package main

import (
	"fmt"
	"github.com/toshi3221/theta_v2"
)

func main() {

	host := "http://dummy-host"

	client, _ := theta_v2.NewClient(host)
	_, error := client.Info()

	if error != nil {
		fmt.Println("Error:", error.Error())
	}
}
