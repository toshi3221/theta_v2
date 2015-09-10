package command

type ListImagesParameters struct {
	EntryCount        *int    `json:"entryCount"`
	MaxSize           *int    `json:"maxSize,omitempty"`
	ContinuationToken *string `json:"continuationToken,omitempty"`
	IncludeThumb      *bool   `json:"includeThumb,omitempty"`
}
type ImageEntry struct {
	Name         *string
	Uri          *string
	Size         *int
	DateTimeZone *string
	Lat          *float32
	Lng          *float32
	Width        *int
	Height       *int
	Thumbnail    *string
}
type ListImagesResults struct {
	Entries           *[]ImageEntry
	TotalEntries      *int
	ContinuationToken *string
}
type ListImagesCommand struct {
	Parameters ListImagesParameters
	Results    ListImagesResults
}

func (command *ListImagesCommand) GetName() string {
	return `camera.listImages`
}
func (command *ListImagesCommand) GetParameters() interface{} {
	return &command.Parameters
}
func (command *ListImagesCommand) GetResults() interface{} {
	return &command.Results
}
