package benchmarks

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"

	"github.com/goph/logur"
	"github.com/goph/logur/adapters/logrusadapter"
)

func newLogrus() logur.Logger {
	logger := logrus.New()
	logger.Level = logrus.TraceLevel
	logger.Out = ioutil.Discard

	return logrusadapter.New(logger)
}

func newDisabledLogrus() logur.Logger {
	logger := logrus.New()
	logger.Level = logrus.ErrorLevel
	logger.Out = ioutil.Discard

	return logrusadapter.New(logger)
}
