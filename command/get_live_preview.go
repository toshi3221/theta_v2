package command

type GetLivePreviewParameters struct {
	SessionId *string `json:"sessionId"`
}
type GetLivePreviewCommand struct {
	Parameters GetLivePreviewParameters
}

func (command *GetLivePreviewCommand) GetName() string {
	return `camera.closeSession`
}
func (command *GetLivePreviewCommand) GetParameters() interface{} {
	return &command.Parameters
}
func (command *GetLivePreviewCommand) GetResults() interface{} {
	return nil
}
