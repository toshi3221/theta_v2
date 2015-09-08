package command

type UpdateSessionParameters struct {
	SessionId string `json:"sessionId"`
	Timeout   int    `json:"timeout,omitempty"`
}
type UpdateSessionResults struct {
	SessionId string
	Timeout   int
}
type UpdateSessionCommand struct {
	Parameters UpdateSessionParameters
	Results    UpdateSessionResults
}

func (command *UpdateSessionCommand) GetName() string {
	return `camera.updateSession`
}
func (command *UpdateSessionCommand) GetParameters() interface{} {
	return &command.Parameters
}
func (command *UpdateSessionCommand) GetResults() interface{} {
	return &command.Results
}
