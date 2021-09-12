package output_processor

type OutputProcessor interface {
	ProcessOutput(interface{})
	ProcessError(err error)
}
