package benchmarks

import (
	"io/ioutil"

	"github.com/hashicorp/go-hclog"

	"github.com/goph/logur"
	"github.com/goph/logur/adapters/hclogadapter"
)

func newHclog() logur.Logger {
	logger := hclog.New(&hclog.LoggerOptions{
		Level:  hclog.Trace,
		Output: ioutil.Discard,
	})

	return hclogadapter.New(logger)
}

func newDisabledHclog() logur.Logger {
	logger := hclog.New(&hclog.LoggerOptions{
		Level:  hclog.Error,
		Output: ioutil.Discard,
	})

	return hclogadapter.New(logger)
}
