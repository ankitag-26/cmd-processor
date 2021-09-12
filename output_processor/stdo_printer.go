package output_processor

import "fmt"

type StandardOutputPrinter struct {
}

func GetStdOutputPrinter() OutputProcessor {
	return StandardOutputPrinter{}
}

func (this StandardOutputPrinter) ProcessOutput(obj interface{}) {
	fmt.Println(obj)
}

func (this StandardOutputPrinter) ProcessError(err error) {
	fmt.Errorf("error ocurred during command execution : %s", err.Error())
}
