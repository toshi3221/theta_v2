package main

import (
	"fmt"
	"github.com/toshi3221/theta_v2"
	"github.com/toshi3221/theta_v2/command"
	"os"
	"time"
)

func main() {

	host := "http://192.168.1.1"
	if len(os.Args) > 1 {
		host = os.Args[1]
	}

	client, _ := theta_v2.NewClient(host)

	// camera.startSession
	fmt.Println("camera.startSession:")
	startSessionCommand := new(command.StartSessionCommand)
	client.CommandExecute(startSessionCommand)

	fmt.Println("  sessionId:", *startSessionCommand.Results.SessionId)
	sessionId := startSessionCommand.Results.SessionId

	// state
	state, _ := client.State()
	fingerprint := *state.Fingerprint

	// camera.takePicture
	fmt.Println("camera.takePicture")
	takePictureCommand := new(command.TakePictureCommand)
	takePictureCommand.Parameters.SessionId = sessionId
	takePictureResponse, _ := client.CommandExecute(takePictureCommand)

	fmt.Println("  state:", *takePictureResponse.State)
	fmt.Println("  commandId:", *takePictureResponse.Id)
	commandId := takePictureResponse.Id

	// Wait Saving File
	fmt.Print("Wait Saving File...")
	lastFingerprint := fingerprint
	for lastFingerprint == fingerprint {
		lastFingerprint = fingerprint
		checkForUpdates, _ := client.CheckForUpdates(lastFingerprint, nil)
		fingerprint = *checkForUpdates.StateFingerprint
	}
	fmt.Println("done.")

	// camera.updateSession
	fmt.Println("camera.updateaSession:")
	updateSessionCommand := new(command.UpdateSessionCommand)
	updateSessionCommand.Parameters.SessionId = sessionId
	updateSessionResponse, _ := client.CommandExecute(updateSessionCommand)

	fmt.Println("  state:", *updateSessionResponse.State)
	sessionId = updateSessionCommand.Results.SessionId

	// camera.closeSession
	closeSessionCommand := new(command.CloseSessionCommand)
	closeSessionCommand.Parameters.SessionId = sessionId
	closeSessionResponse, _ := client.CommandExecute(closeSessionCommand)

	fmt.Println("camera.closeSession:")
	fmt.Println("  state:", *closeSessionResponse.State)

	// CommandStatus
	fmt.Println("CommandStatus:")
	commandStatusResponse, _ := client.CommandStatus(*commandId, takePictureCommand)
	for *commandStatusResponse.State != "done" {
		time.Sleep(200 * time.Millisecond)
		commandStatusResponse, _ = client.CommandStatus(*commandId, takePictureCommand)
	}
	fmt.Println("  state:", *commandStatusResponse.State)
	fmt.Println("  fileUri:", *takePictureCommand.Results.FileUri)
}
