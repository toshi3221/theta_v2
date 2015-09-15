package command

type StartSessionParameters struct {
}
type StartSessionResults struct {
	SessionId *string
	Timeout   *int
}
type StartSessionCommand struct {
	Parameters StartSessionParameters
	Results    StartSessionResults
}

func (command *StartSessionCommand) GetName() string {
	return `camera.startSession`
}
func (command *StartSessionCommand) GetParameters() interface{} {
	return &command.Parameters
}
func (command *StartSessionCommand) GetResults() interface{} {
	return &command.Results
}
