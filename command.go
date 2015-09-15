package theta_v2

type Command interface {
	GetName() string
	GetParameters() interface{}
	GetResults() interface{}
}
