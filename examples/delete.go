package main

import (
	"fmt"
	"github.com/toshi3221/osc"
	"github.com/toshi3221/osc/command"
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

	client, _ := osc.NewClient(host)

	deleteCommand := new(command.DeleteCommand)
	deleteCommand.Parameters.FileUri = &fileUri
	response, _ := client.CommandExecute(deleteCommand)
	fmt.Println("state:", *response.State)
}
