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

	fmt.Println("camera.startSession:")
	fmt.Println("  sessionId:", *startSessionCommand.Results.SessionId)
	sessionId := startSessionCommand.Results.SessionId

	// camera.takePicture
	takePictureCommand := new(command.TakePictureCommand)
	takePictureCommand.Parameters.SessionId = sessionId
	takePictureResponse, _ := client.CommandExecute(takePictureCommand)

	fmt.Println("camera.takePicture")
	fmt.Println("  state:", *takePictureResponse.State)
	fmt.Println("  commandId:", *takePictureResponse.Id)
	commandId := takePictureResponse.Id

	// camera.updateSession
	updateSessionCommand := new(command.UpdateSessionCommand)
	updateSessionCommand.Parameters.SessionId = sessionId
	updateSessionResponse, _ := client.CommandExecute(updateSessionCommand)

	fmt.Println("camera.updateaSession:")
	fmt.Println("  state:", *updateSessionResponse.State)
	sessionId = updateSessionCommand.Results.SessionId

	// camera.closeSession
	closeSessionCommand := new(command.CloseSessionCommand)
	closeSessionCommand.Parameters.SessionId = sessionId
	closeSessionResponse, _ := client.CommandExecute(closeSessionCommand)

	fmt.Println("camera.closeSession:")
	fmt.Println("  state:", *closeSessionResponse.State)

	// CommandStatus
	commandStatusResponse, _ := client.CommandStatus(*commandId, takePictureCommand)
	fmt.Println("CommandStatus:")
	fmt.Println("  state:", *commandStatusResponse.State)
	fmt.Println("  fileUri:", *takePictureCommand.Results.FileUri)
}
