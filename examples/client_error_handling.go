package main

import (
	"fmt"
	"github.com/toshi3221/osc"
)

func main() {

	host := "http://dummy-host"

	client, _ := osc.NewClient(host)
	_ , error := client.Info()

	if error != nil {
		fmt.Println("Error:", error.Error())
	}
}
