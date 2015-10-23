package command

type ListAllParameters struct {
	EntryCount        *int    `json:"entryCount"`
	ContinuationToken *string `json:"continuationToken,omitempty"`
	Detail            *bool   `json:"detail,omitempty"`
	Sort              *string `json:"sort,omitempty"`
}
type ListAllEntry struct {
	Name                   *string  `json:"name,omitempty"`
	Uri                    *string  `json:"uri,omitempty"`
	Size                   *int     `json:"size,omitempty"`
	DateTimeZone           *string  `json:"dateTimeZone,omitempty"`
	DateTime               *string  `json:"dateTime,omitempty"`
	Lat                    *float32 `json:"lat,omitempty"`
	Lng                    *float32 `json:"lng,omitempty"`
	Width                  *int     `json:"width,omitempty"`
	Height                 *int     `json:"height,omitempty"`
	IntervalCaptureGroupId *string  `json:"_intervalCaptureGroupId,omitempty"`
	RecordTime             *int     `json:"recordTime,omitempty"`
}
type ListAllResults struct {
	Entries           *[]ListAllEntry `json:"entries,omitempty"`
	TotalEntries      *int            `json:"totalEntries,omitempty"`
	ContinuationToken *string         `json:"continuationToken,omitempty"`
}
type ListAllCommand struct {
	Parameters ListAllParameters
	Results    ListAllResults
}

func (command *ListAllCommand) GetName() string {
	return `camera._listAll`
}
func (command *ListAllCommand) GetParameters() interface{} {
	return &command.Parameters
}
func (command *ListAllCommand) GetResults() interface{} {
	return &command.Results
}
