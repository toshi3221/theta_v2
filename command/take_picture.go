package command

type TakePictureParameters struct {
	SessionId *string `json:"sessionId"`
}
type TakePictureResults struct {
	FileUri *string
}
type TakePictureCommand struct {
	Parameters TakePictureParameters
	Results    TakePictureResults
}

func (command *TakePictureCommand) GetName() string {
	return `camera.takePicture`
}
func (command *TakePictureCommand) GetParameters() interface{} {
	return &command.Parameters
}
func (command *TakePictureCommand) GetResults() interface{} {
	return &command.Results
}
