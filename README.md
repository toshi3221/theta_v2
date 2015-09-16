## theta_v2
RICOH THETA API v2 Client for Golang

## Example
If you know other API usage, see and execute [examples](examples) go source. 
```sh
$ cd $GOPATH/src/github.com/toshi3221/theta_v2/examples
$ go run info.go http://(theta_v2-host)
```

### camera.listImages Command
```go
package main

import (
  "fmt"
  "github.com/toshi3221/theta_v2"
  "github.com/toshi3221/theta_v2/command"
)

func main() {

  client, _ := theta_v2.NewClient("host")

  command := new(command.ListImagesCommand)
  parameters := &command.Parameters

  entryCount, maxSize, includeThumb := 10, 10, false
  parameters.EntryCount = &entryCount
  parameters.MaxSize = &maxSize
  parameters.IncludeThumb = &includeThumb

  client.CommandExecute(command)

  results := command.Results
  if *results.TotalEntries > 0 {
    entries := *results.Entries
    for i := range entries {
      fmt.Printf("{name: %s, uri: %s}\n", *entries[i].Name, *entries[i].Uri)
    }
  }

}
```

## The MIT License (MIT)

Copyright (c) 2015 Toshiharu Takematsu

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
