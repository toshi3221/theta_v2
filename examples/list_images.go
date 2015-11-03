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

	command := new(command.ListImagesCommand)
	parameters := &command.Parameters

	entryCount, includeThumb := 10, false
	parameters.EntryCount = &entryCount
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
			fmt.Printf("{name: %s, uri: %s, size: %d}\n", *entries[i].Name, *entries[i].Uri, *entries[i].Size)
		}
	}

}
