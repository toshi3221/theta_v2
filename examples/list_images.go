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

	client, _ := osc.NewClient(host)

	listImagesCommand := new(command.ListImagesCommand)
	entryCount, maxSize, includeThumb := 10, 10, true
	listImagesCommand.Parameters.EntryCount = &entryCount
	listImagesCommand.Parameters.MaxSize = &maxSize
	listImagesCommand.Parameters.IncludeThumb = &includeThumb
	client.CommandExecute(listImagesCommand)

	results := listImagesCommand.Results
	fmt.Println("totalEntries:", *results.TotalEntries)
	if *results.TotalEntries > 0 {
		entry := (*results.Entries)[0]
		fmt.Println("entries[0].name:", *entry.Name)
		fmt.Println("entries[0].uri:", *entry.Uri)
	}

}
