package command

type GetMetadataParameters struct {
	FileUri string `json:"fileUri"`
}
type Exif struct {
	ExifVersion         string
	ImageDescription    string
	DateTime            string
	ImageWidth          int
	ImageLength         int
	ColorSpace          int
	Compression         int
	Orientation         int
	Flash               int
	FocalLength         float32
	WhiteBalance        int
	ExposureTime        float32
	FNumber             float32
	ExposureProgram     int
	ISOSpeedRatings     int
	ApertureValue       float32
	BrightnessValue     int
	ExposureBiasValue   int
	GPSProcessingMethod int
	GPSLatitudeRef      string
	GPSLatitude         float32
	GPSLongitudeRef     string
	GPSLongitude        float32
	Make                string
	Model               string
	Software            string
	Copyright           string
	MakerNote           string
}
type Xmp struct {
	ProjectionType               string
	UsePanoramaViewer            bool
	PoseHeadingDegrees           float32
	CroppedAreaImageWidthPixels  int
	CroppedAreaImageHeightPixels int
	FullPanoWidthPixels          int
	FullPanoHeightPixels         int
	CroppedAreaLeftPixels        int
	CroppedAreaTopPixels         int
}
type GetMetadataResults struct {
	Exif Exif
	Xmp  Xmp
}
type GetMetadataCommand struct {
	Parameters GetMetadataParameters
	Results    GetMetadataResults
}

func (command *GetMetadataCommand) GetName() string {
	return `camera.getMetadata`
}
func (command *GetMetadataCommand) GetParameters() interface{} {
	return &command.Parameters
}
func (command *GetMetadataCommand) GetResults() interface{} {
	return &command.Results
}
