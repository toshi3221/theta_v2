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

	client, _ := theta_v2.NewClient(host)

	// camera.startSession
	startSessionCommand := new(command.StartSessionCommand)
	client.CommandExecute(startSessionCommand)
	sessionId := startSessionCommand.Results.SessionId

	// camera.setOptions
	fmt.Println("camera.setOptions:")
	setOptionsCommand := new(command.SetOptionsCommand)
	setOptionsCommand.Parameters.SessionId = sessionId
	setOptions := new(command.Options)

	// RICOH THETA S is NOT supported 5000x2500 jpeg
	image_type, image_width, image_height := "jpeg", 5000, 2000

	setOptions.FileFormat = &command.FileFormat{Type: &image_type, Width: &image_width, Height: &image_height}
	setOptionsCommand.Parameters.Options = setOptions
	setOptionsResponse, _ := client.CommandExecute(setOptionsCommand)

	// Error Handling
	state := *setOptionsResponse.State
	fmt.Println("  State:", state)
	switch state {
	case "error":
		fmt.Println("  HTTP Response:")
		fmt.Println("    Status:", client.Response.Status)
		fmt.Println("    StatusCode:", client.Response.StatusCode)
		fmt.Println("  Error:")
		fmt.Println("    Code:", *setOptionsResponse.Error.Code)
		fmt.Println("    Message:", *setOptionsResponse.Error.Message)
	}

	// camera.closeSession
	closeSessionCommand := new(command.CloseSessionCommand)
	closeSessionCommand.Parameters.SessionId = sessionId
	client.CommandExecute(closeSessionCommand)

}
