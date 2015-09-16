package command

type FinishWlanParameters struct {
	SessionId *string `json:"sessionId"`
}
type FinishWlanResults struct {
}
type FinishWlanCommand struct {
	Parameters FinishWlanParameters
	Results    FinishWlanResults
}

func (command *FinishWlanCommand) GetName() string {
	return `camera._finishWlan`
}
func (command *FinishWlanCommand) GetParameters() interface{} {
	return &command.Parameters
}
func (command *FinishWlanCommand) GetResults() interface{} {
	return &command.Results
}
