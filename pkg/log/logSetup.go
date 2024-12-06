package log

import (
	"github.com/CodingCat12/pacgraph/pkg/config"
	"github.com/sirupsen/logrus"
)

var Logger = logrus.New()

func LoggerSetup() {
	if config.AdjustedConfig.DebugMode {
		Logger.SetLevel(logrus.DebugLevel)
	} else {
		Logger.SetLevel(logrus.WarnLevel)
	}
}
