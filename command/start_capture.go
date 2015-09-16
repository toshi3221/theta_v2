package command

type StartCaptureParameters struct {
	SessionId *string `json:"sessionId"`
}
type StartCaptureResults struct {
}
type StartCaptureCommand struct {
	Parameters StartCaptureParameters
	Results    StartCaptureResults
}

func (command *StartCaptureCommand) GetName() string {
	return `camera._startCapture`
}
func (command *StartCaptureCommand) GetParameters() interface{} {
	return &command.Parameters
}
func (command *StartCaptureCommand) GetResults() interface{} {
	return &command.Results
}
