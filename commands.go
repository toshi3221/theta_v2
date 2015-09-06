package osc

type Command interface {
	GetName() (string)
	GetParameters() (interface{})
	GetResults() (interface{})
}


type StartSessionParameters struct {
        WaitTimeout int `json:"waitTimout,omitempty"`
}
type StartSessionResults struct {
        SessionId string
        Timeout   int
}
type StartSessionCommand struct {
	Parameters	StartSessionParameters
	Results		StartSessionResults
}
func (command *StartSessionCommand) GetName() (string) {
	return `camera.startSession`
}
func (command *StartSessionCommand) GetParameters() (interface{}) {
	return &command.Parameters
}
func (command *StartSessionCommand) GetResults() (interface{}) {
	return &command.Results
}


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


type TakePictureParameters struct {
        SessionId string `json:"sessionId"`
}
type TakePictureResults struct {
        FileUri string
}
type TakePictureCommand struct {
	Parameters	TakePictureParameters
	Results		TakePictureResults
}
func (command *TakePictureCommand) GetName() (string) {
	return `camera.takePicture`
}
func (command *TakePictureCommand) GetParameters() (interface{}) {
	return &command.Parameters
}
func (command *TakePictureCommand) GetResults() (interface{}) {
	return &command.Results
}

