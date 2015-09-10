package command

type DeleteParameters struct {
	FileUri *string `json:"fileUri"`
}
type DeleteResults struct {
}
type DeleteCommand struct {
	Parameters DeleteParameters
	Results    DeleteResults
}

func (command *DeleteCommand) GetName() string {
	return `camera.getImage`
}
func (command *DeleteCommand) GetParameters() interface{} {
	return &command.Parameters
}
func (command *DeleteCommand) GetResults() interface{} {
	return &command.Results
}
