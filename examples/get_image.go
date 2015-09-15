package main

import (
	"github.com/toshi3221/theta_v2"
	"github.com/toshi3221/theta_v2/command"
	"io"
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

	getImageCommand := new(command.GetImageCommand)
	getImageCommand.Parameters.FileUri = &fileUri
	response, _ := client.CommandExecute(getImageCommand)
	http_body := response.Results.(io.ReadCloser)
	defer http_body.Close()
	out, _ := os.Create(fileUri)
	defer out.Close()
	io.Copy(out, http_body)
}
