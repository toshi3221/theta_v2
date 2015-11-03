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

	command := new(command.ListImagesCommand)
	parameters := &command.Parameters

	entryCount, maxSize, includeThumb := 10, 10, true
	parameters.EntryCount = &entryCount
	parameters.MaxSize = &maxSize
	parameters.IncludeThumb = &includeThumb

	_, error := client.CommandExecute(command)
	if error != nil {
		fmt.Println("Error:", error)
		return
	}

	results := command.Results
	fmt.Println("totalEntries:", *results.TotalEntries)
	if *results.TotalEntries > 0 {
		entries := *results.Entries
		for i := range entries {
			fmt.Printf("{name: %s, uri: %s}\n", *entries[i].Name, *entries[i].Uri)
		}
	}

}
