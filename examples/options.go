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

	// camera.getOptions
	getOptionsCommand := new(command.GetOptionsCommand)
	getOptionsCommand.Parameters.SessionId = sessionId
	optionNames := &[]string{"fileFormat", "fileFormatSupport", "captureMode", "captureModeSupport", "exposureProgram", "exposureProgramSupport", "shutterSpeed", "iso"}
	getOptionsCommand.Parameters.OptionNames = optionNames
	client.CommandExecute(getOptionsCommand)

	options := getOptionsCommand.Results.Options
	fmt.Println("camera.getOptions:")
	fmt.Println("  exposureProgram:", *options.ExposureProgram)
	fmt.Printf("  fileFormat: {type: %s, width: %d, height: %d}\n", *options.FileFormat.Type, *options.FileFormat.Width, *options.FileFormat.Height)
	fmt.Println("  captureMode:", *options.CaptureMode)
	fmt.Println("  ISO:", *options.Iso)

	// camera.setOptions
	setOptionsCommand := new(command.SetOptionsCommand)
	setOptionsCommand.Parameters.SessionId = sessionId
	setOptions := new(command.Options)
	image_type, image_width, image_height := "jpeg", 5376, 2688
	setOptions.FileFormat = &command.FileFormat{Type: &image_type, Width: &image_width, Height: &image_height}
	iso := 100
	setOptions.Iso = &iso
	exposureProgram := 9 // ISO priority program
	setOptions.ExposureProgram = &exposureProgram
	setOptionsCommand.Parameters.Options = setOptions
	setOptionsResponse, _ := client.CommandExecute(setOptionsCommand)

	fmt.Println("camera.setOptions:")
	fmt.Println("  State:", *setOptionsResponse.State)

	// camera.closeSession
	closeSessionCommand := new(command.CloseSessionCommand)
	closeSessionCommand.Parameters.SessionId = sessionId
	closeSessionResponse, _ := client.CommandExecute(closeSessionCommand)

	fmt.Println("camera.closeSession:")
	fmt.Println("  State:", *closeSessionResponse.State)

}
