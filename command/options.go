package command

type FileFormat struct {
	Type   *string `json:"type"`
	Width  *int    `json:"width"`
	Height *int    `json:"height"`
}
type GpsInfo struct {
	Lat *float32 `json:"lat"`
	Lng *float32 `json:"lng"`
}
type ExposureBracket struct {
	Shots     *int     `json:"shots,omitempty"`
	Increment *float32 `json:"increment,omitempty"`
	AutoMode  *bool    `json:"autoMode,omitempty"`
}
type ExposureBracketSupport struct {
	ShotsSupport     *[]int     `json:"shotsSupport,omitempty"`
	IncrementSupport *[]float32 `json:"incrementSupport,omitempty"`
	AutoMode         *bool      `json:"autoMode,omitempty,bool"`
}
type Options struct {
	CaptureMode                 *string                 `json:"captureMode,omitempty"`
	CaptureModeSupport          *[]string               `json:"captureModeSupport,omitempty"`
	ExposureProgram             *int                    `json:"exposureProgram,omitempty,int"`
	ExposureProgramSupport      *[]int                  `json:"exposureProgramSupport,omitempty"`
	Iso                         *int                    `json:"iso,omitempty"`
	IsoSupport                  *[]int                  `json:"isoSupport,omitempty"`
	ShutterSpeed                *float32                `json:"shutterSpeed,omitempty,float"`
	ShutterSpeedSupport         *[]float32              `json:"shutterSpeedSupport,omitempty"`
	Aperture                    *float32                `json:"aperture,omitempty,float"`
	ApertureSupport             *[]float32              `json:"apertureSupport,omitempty"`
	WhiteBalance                *string                 `json:"whiteBalance,omitempty,string"`
	WhiteBalanceSupport         *[]string               `json:"whiteBalanceSupport,omitempty"`
	ExposureCompensation        *float32                `json:"exposureCompensation,omitempty,float"`
	ExposureCompensationSupport *[]float32              `json:"exposureCompensationSupport,omitempty"`
	FileFormat                  *FileFormat             `json:"fileFormat,omitempty,FileFormat"`
	FileFormatSupport           *[]FileFormat           `json:"fileFormatSupport,omitempty"`
	ExposureDelay               *int                    `json:"exposureDelay,omitempty"`
	ExposureDelaySupport        *[]int                  `json:"exposureDelaySupport,omitempty"`
	SleepDelay                  *int                    `json:"sleepDelay,omitempty"`
	SleepDelaySupport           *[]int                  `json:"sleepDelaySupport,omitempty"`
	OffDelay                    *int                    `json:"offDelay,omitempty"`
	OffDelaySupport             *[]int                  `json:"offDelaySupport,omitempty"`
	TotalSpace                  *int                    `json:"totalSpace,omitempty"`
	RemainingSpace              *int                    `json:"remainingSpace,omitempty"`
	RemainingPictures           *int                    `json:"remainingPictures,omitempty"`
	GpsInfo                     *GpsInfo                `json:"gpsInfo,omitempty"`
	DateTimeZone                *string                 `json:"dateTimeZone,omitempty"`
	Hdr                         *bool                   `json:"hdr,omitempty"`
	HdrSupport                  *bool                   `json:"hdrSupport,omitempty"`
	ExposureBracket             *ExposureBracket        `json:"exposureBracket,omitempty"`
	ExposureBracketSupport      *ExposureBracketSupport `json:"exposureBracketSupport,omitempty"`
	Gyro                        *bool                   `json:"gyro,omitempty"`
	GyroSupport                 *bool                   `json:"gyroSupport,omitempty"`
	Gps                         *bool                   `json:"gps,omitempty"`
	GpsSupport                  *bool                   `json:"gpsSupport,omitempty"`
	ImageStabilization          *string                 `json:"imageStabilization,omitempty"`
	ImageStabilizationSupport   *[]string               `json:"imageStabilizationSupport,omitempty"`
	WifiPassword                *string                 `json:"wifiPassword,omitempty"`
}

type GetOptionsParameters struct {
	SessionId   *string   `json:"sessionId"`
	OptionNames *[]string `json:"optionNames"`
}
type GetOptionsResults struct {
	Options *Options
}
type GetOptionsCommand struct {
	Parameters GetOptionsParameters
	Results    GetOptionsResults
}

func (command *GetOptionsCommand) GetName() string {
	return `camera.getOptions`
}
func (command *GetOptionsCommand) GetParameters() interface{} {
	return &command.Parameters
}
func (command *GetOptionsCommand) GetResults() interface{} {
	return &command.Results
}

type SetOptionsParameters struct {
	SessionId *string  `json:"sessionId"`
	Options   *Options `json:"options"`
}
type SetOptionsResults struct {
}
type SetOptionsCommand struct {
	Parameters SetOptionsParameters
	Results    SetOptionsResults
}

func (command *SetOptionsCommand) GetName() string {
	return `camera.setOptions`
}
func (command *SetOptionsCommand) GetParameters() interface{} {
	return &command.Parameters
}
func (command *SetOptionsCommand) GetResults() interface{} {
	return &command.Results
}
