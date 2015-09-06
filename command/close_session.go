package command

type CloseSessionParameters struct {
        SessionId string `json:"sessionId"`
}
type CloseSessionResults struct {
}
type CloseSessionCommand struct {
	Parameters	CloseSessionParameters
	Results		CloseSessionResults
}
func (command *CloseSessionCommand) GetName() (string) {
	return `camera.closeSession`
}
func (command *CloseSessionCommand) GetParameters() (interface{}) {
	return &command.Parameters
}
func (command *CloseSessionCommand) GetResults() (interface{}) {
	return &command.Results
}
