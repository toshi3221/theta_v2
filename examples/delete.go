package main

import (
	"fmt"
	"github.com/toshi3221/theta_v2"
	"github.com/toshi3221/theta_v2/command"
	"os"
)

func main() {

	host := "http://192.168.1.1"
	if len(os.Args) > 1 {
		host = os.Args[1]
	}
	fileUri := "1"
	if len(os.Args) > 2 {
		fileUri = os.Args[2]
	}

	client, _ := theta_v2.NewClient(host)

	deleteCommand := new(command.DeleteCommand)
	deleteCommand.Parameters.FileUri = &fileUri
	response, _ := client.CommandExecute(deleteCommand)
	fmt.Println("state:", *response.State)
}
