package command

type GetImageParameters struct {
	FileUri string `json:"fileUri"`
	MaxSize int    `json:"maxSize,omitempty"`
}
type GetImageCommand struct {
	Parameters GetImageParameters
}

func (command *GetImageCommand) GetName() string {
	return `camera.getImage`
}
func (command *GetImageCommand) GetParameters() interface{} {
	return &command.Parameters
}
func (command *GetImageCommand) GetResults() interface{} {
	return nil
}
