package osc

type Command interface {
	GetName() string
	GetParameters() interface{}
	GetResults() interface{}
}
