package command

type StopCaptureParameters struct {
	SessionId *string `json:"sessionId"`
}
type StopCaptureResults struct {
}
type StopCaptureCommand struct {
	Parameters StopCaptureParameters
	Results    StopCaptureResults
}

func (command *StopCaptureCommand) GetName() string {
	return `camera._stopCapture`
}
func (command *StopCaptureCommand) GetParameters() interface{} {
	return &command.Parameters
}
func (command *StopCaptureCommand) GetResults() interface{} {
	return &command.Results
}
