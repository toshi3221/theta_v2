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

	getMetadataCommand := new(command.GetMetadataCommand)
	getMetadataCommand.Parameters.FileUri = &fileUri
	client.CommandExecute(getMetadataCommand)

	exif := getMetadataCommand.Results.Exif
	fmt.Println("Exif:")
	fmt.Println("  ExifVersion:", *exif.ExifVersion)
	fmt.Println("  ImageDescription:", *exif.ImageDescription)
	xmp := getMetadataCommand.Results.Xmp
	fmt.Println("XMP:")
	fmt.Println("  ProjectionType:", *xmp.ProjectionType)
	fmt.Println("  UsePanoramaViewer:", *xmp.UsePanoramaViewer)

}
